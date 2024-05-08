package user_v1

import (
	"apps/auth/utils"
)

type UserController struct {
	logger utils.Logger
}


func NewUserController(logger utils.Logger) UserController {
	return UserController{
		logger: logger,
	}
}