package main

import (
    "github.com/joho/godotenv"
	"go.uber.org/fx"
	"apps/auth/utils"
	"apps/auth/bootstrap"
)

func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
