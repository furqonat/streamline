package auth_v1

import (
	"apps/auth/dto"
	"apps/auth/services"
	"apps/auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger  utils.Logger
	service services.AuthService
}

func NewAuthController(logger utils.Logger, service services.AuthService) AuthController {
	return AuthController{
		logger:  logger,
		service: service,
	}
}

func (a AuthController) SignIn(c *gin.Context) {
	data := dto.SignInDto{}

	if err := c.ShouldBindJSON(&data); err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}

	token, err := a.service.SignIn(data)
	if err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign in successful", "token": *token})
}

func (a AuthController) SignUp(c *gin.Context) {
	data := dto.SignUpDto{}

	if err := c.ShouldBindJSON(&data); err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}

	token, err := a.service.SignUp(data)
	if err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign up successful", "token": *token})
}
