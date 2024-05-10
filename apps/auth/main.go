package main

import (
    "github.com/joho/godotenv"
	"go.uber.org/fx"
	"fmt"
	"apps/auth/utils"
	"apps/auth/bootstrap"
)

func main() {
	if err := godotenv.Load(); err != nil {
	  fmt.Printf("unable load env")
	}
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
