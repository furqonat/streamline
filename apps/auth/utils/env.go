package utils

import (
	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	JwtSecret   string `mapstructure:"JWT_SECRET"`
}

// NewEnv creates a new environment
func NewEnv(log Logger) Env {

	env := Env{}
	viper.SetConfigFile(".env")
	viper.SetConfigFile("../../../.env") // for testing purpose
	viper.SetConfigType("env")
	viper.SetConfigFile("../../.env") // for development purpose
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Infof("%+v \n", env)
	return env
}
