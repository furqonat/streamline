package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewMiscRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewUserRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	miscRoutes MiscRoutes,
	authRoutes AuthRoutes,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		miscRoutes,
		authRoutes,
		userRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
