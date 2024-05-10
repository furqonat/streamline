package routes

import (
	auth_v1 "apps/auth/controllers/v1/auth"
	"apps/auth/utils"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	authController auth_v1.AuthController
}

// Setup Misc routes
func (s AuthRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/apis/v1")
	{
		api.POST("/signin", s.authController.SignIn)
		api.POST("/signup", s.authController.SignUp)
	}
}

// NewAuthRoutes creates new Misc controller
func NewAuthRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	authController auth_v1.AuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
