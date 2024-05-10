package controllers

import (
	auth_v1 "apps/auth/controllers/v1/auth"
	misc_v1 "apps/auth/controllers/v1/misc"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(misc_v1.NewMiscController),
	fx.Provide(auth_v1.NewAuthController),
)
