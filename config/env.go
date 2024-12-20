package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName                 string
	DBPassword             string
	DBUsername             string
	DBPort                 string
	DBHost                 string
	DBSchema               string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBName:                 getEnv("DB_DATABASE", "habits"),
		DBPassword:             getEnv("DB_PASSWORD", "password"),
		DBUsername:             getEnv("DB_USERNAME", "username"),
		DBPort:                 getEnv("DB_PORT", "8080"),
		DBHost:                 getEnv("DB_HOST", "http://localhost"),
		DBSchema:               getEnv("DB_SCHEMA", "public"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
