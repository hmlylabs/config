package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type Config struct {
	ServerName string `mapstructure:"SERVER_NAME"`
	Port       string `mapstructure:"SERVER_PORT"`
	DbConfig   DatabaseConfig
}

func LoadConfig() (cfg Config, err error) {
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
