package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBTimezone string
	AppConfig string
	GinPort string
	GinHost string
	RedisHost string
	RedisPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("env don't loading")
	}
	return &Config{
		DBHost: getEnv("POSTGRES_HOST"),
		DBPort: getEnv("POSTGRES_PORT"),
		DBUser: getEnv("POSTGRES_USER"),
		DBPass: getEnv("POSTGRES_PASS"),
		DBName: getEnv("POSTGRES_DB"),
		DBTimezone: getEnv("POSTGRES_TIMEZONE"),
		AppConfig: getEnv("APP_ENV"),
		GinPort: getEnv("GIN_PORT"),
		GinHost: getEnv("GIN_HOST"),
		RedisHost: getEnv("REDIS_HOST"),
		RedisPort: getEnv("REDIS_PORT"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Missing env variable: %s\n", key)
	}
	return value
}