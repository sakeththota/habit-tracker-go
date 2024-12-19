package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DBPassword string
	DBUsername string
	DBPort     string
	DBHost     string
	DBSchema   string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBName:     getEnv("DB_DATABASE", "habits"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBUsername: getEnv("DB_USERNAME", "username"),
		DBPort:     getEnv("DB_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "http://localhost"),
		DBSchema:   getEnv("DB_SCHEMA", "public"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
