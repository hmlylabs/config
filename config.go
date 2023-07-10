package config

import (
	"os"
)

type DatabaseConfig struct {
	Port     string
	Name     string
	Host     string
	User     string
	Password string
}

type Config struct {
	ServerName string
	Port       string
	Database   DatabaseConfig
}

func New() *Config {
	return &Config{
		ServerName: getEnv("SERVER_NAME", "meal-server"),
		Port:       getEnv("PORT", "8080"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "meal_ad"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "meal_db"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
