package user_v1

import (
	"apps/auth/dto"
	"apps/auth/services"
	"apps/auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	logger  utils.Logger
	service services.UserService
}

func NewUserController(logger utils.Logger, service services.UserService) UserController {
	return UserController{
		logger:  logger,
		service: service,
	}
}

func (u UserController) GetUser(c *gin.Context) {
	u.logger.Info("Getting user")

	id := c.GetString(utils.UID)
	if id == "" {
		c.JSON(http.StatusUnauthorized, utils.ResponseError{Message: "User ID not found"})
		return
	}
	user, err := u.service.GetUser(id)
	if err != nil {
		u.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOk{Data: user, Message: "User fetched successfully"})
}

func (u UserController) UpdateUser(c *gin.Context) {
	u.logger.Info("Updating user")

	id := c.GetString(utils.UID)
	json := dto.UpdateUserDto{}

	if err := c.ShouldBindJSON(&json); err != nil {
		u.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	user, err := u.service.UpdateUser(id, json)
	if err != nil {
		u.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOk{Data: user, Message: "User updated successfully"})
}
