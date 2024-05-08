package auth_v1

import (
	"apps/auth/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger utils.Logger
}

func NewAuthController(logger utils.Logger) AuthController {
	return AuthController{
		logger: logger,
	}
}

func (a AuthController) SignIn(c *gin.Context) {

}
