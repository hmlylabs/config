package config

import (
	"os"
)

type Config struct {
	ServerName  string
	Port        string
	DatabaseUrl string
	ShadowUrl   string
}

func New() *Config {
	return &Config{
		ServerName:  getEnv("SERVER_NAME", "meal-server"),
		Port:        getEnv("PORT", "8080"),
		DatabaseUrl: getEnv("DATABASE_URL", "localhost:5432/postgres"),
		ShadowUrl:   getEnv("SHADOW_URL", "localhost:5432/postgres"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
