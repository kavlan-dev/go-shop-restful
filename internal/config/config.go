package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerHost string
	ServerPort string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
}

func LoadConfig() (Config, error) {
	config := Config{
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "myuser"),
		DBPassword: getEnv("DB_PASSWORD", "pass"),
		DBName:     getEnv("DB_NAME", "mydb"),
		DBPort:     getEnv("DB_PORT", "5432"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
	}

	if _, err := strconv.Atoi(config.ServerPort); err != nil {
		return Config{}, err
	}
	if _, err := strconv.Atoi(config.DBPort); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetServerAddress(config Config) string {
	return config.ServerHost + ":" + config.ServerPort
}
