package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerName  string `mapstructure:"SERVER_NAME"`
	Port        string `mapstructure:"PORT"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
