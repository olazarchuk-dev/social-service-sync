package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DbUser     string `mapstructure:"db_user"`
	DbPassword string `mapstructure:"db_pass"`
	DbName     string `mapstructure:"db_name"`
	DbHost     string `mapstructure:"db_host"`
}

func LoadConfig() (config Config, err error) {

	mode := os.Getenv("APP_ENV")

	viper.AddConfigPath(".")
	if mode == "prod" {
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigName("env-dev")
	}
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
