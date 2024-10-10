package config

import "os"

type Config struct {
	ServerName  string
	Port        string
	DatabaseUrl string
}

func New() *Config {
	return &Config{
		ServerName:  getEnv("SERVER_NAME", "server"),
		Port:        getEnv("PORT", "8080"),
		DatabaseUrl: getEnv("DatabaseUrl", "localhost"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
