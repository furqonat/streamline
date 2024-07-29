package middlewares

import (
	"apps/auth/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JwtMiddleware struct {
	jwt    utils.Jwt
	logger utils.Logger
}

func NewJwtMiddleware(jwt utils.Jwt, logger utils.Logger) JwtMiddleware {
	return JwtMiddleware{
		jwt:    jwt,
		logger: logger,
	}
}

func (j JwtMiddleware) HandleAuthWithRoles(roles ...string) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		bearerToken, err := j.getTokenFromHeaders(gCtx)
		if err != nil {
			gCtx.JSON(http.StatusUnauthorized, utils.ResponseError{
				Message: "Unauthorized",
			})
			gCtx.Abort()
			return
		}

		token, err := j.jwt.Decode(bearerToken)
		if err != nil {
			j.logger.Info(gin.H{"message": "Unauthorized", "error": err.Error()})
			gCtx.JSON(http.StatusUnauthorized, utils.ResponseError{
				Message: fmt.Sprintf("Error: %s", err.Error()),
			})
			gCtx.Abort()
			return
		}

		if len(roles) < 1 {
			gCtx.JSON(http.StatusInternalServerError, utils.ResponseError{
				Message: "Invalid user roles",
			})
			gCtx.Abort()
			return
		}

		if len(roles) > 0 {
			if ok := j.checkRoleIsValid(roles, token); !ok {
				gCtx.JSON(http.StatusUnauthorized, utils.ResponseError{
					Message: "Unauthorized, invalid roles",
				})
				gCtx.Abort()
				return
			}
		}
		gCtx.Set(utils.UID, token.ID)
		gCtx.Set(utils.ROLES, token.Role)
		gCtx.Next()
	}
}

func (j JwtMiddleware) getTokenFromHeaders(c *gin.Context) (string, error) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		return "", errors.New("no token provided")
	}
	token := strings.TrimPrefix(bearer, "Bearer ")
	return token, nil
}

func (j JwtMiddleware) checkRoleIsValid(roles []string, token *utils.Claims) bool {
	for _, val := range roles {
		for _, key := range token.Role {
			if key.Name == val {
				return true
			}
		}
	}
	return false

}
