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
	// get user agent and ip address
	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	if err := c.ShouldBindJSON(&data); err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}

	data.IpAddress = &ipAddress
	data.UserAgent = &userAgent

	token, err := a.service.SignIn(data)
	if err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOk{Message: "Sign in successful", Data: map[string]string{"token": *token}})
}

func (a AuthController) SignUp(c *gin.Context) {
	data := dto.SignUpDto{}
	// get user agent and ip address
	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	if err := c.ShouldBindJSON(&data); err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}
	data.IpAddress = &ipAddress
	data.UserAgent = &userAgent
	token, err := a.service.SignUp(data)
	if err != nil {
		a.logger.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseOk{Message: "Sign up successful", Data: map[string]string{"token": *token}})
}
