package routes

import (
	user_v1 "apps/auth/controllers/v1/user"
	"apps/auth/middlewares"
	"apps/auth/utils"
)

// UserRoutes struct
type UserRoutes struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	userController user_v1.UserController
	jwtMiddleware  middlewares.JwtMiddleware
}

// Setup Misc routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/apis/v1").Use(s.jwtMiddleware.HandleAuthWithRoles())
	{
		api.GET("/user", s.userController.GetUser)
	}
}

// NewUserRoutes creates new Misc controller
func NewUserRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	userController user_v1.UserController,
	jwtMiddleware middlewares.JwtMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		jwtMiddleware:  jwtMiddleware,
	}
}
