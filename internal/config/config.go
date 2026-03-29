package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Environment string
	Database    databaseConfig
	JWTSecret   string
	CORS        []string
	Admin       adminConfig
}

type databaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
}

type adminConfig struct {
	Name     string
	Password string
	Email    string
}

func InitConfig() (*Config, error) {
	var config Config
	config.Environment = envOrDefault("ENV", "prod")
	config.Database.Host = envOrDefault("DATABASE_HOST", "localhost")
	config.Database.User = envOrDefault("DATABASE_USER", "postgres")
	config.Database.Password = envOrDefault("DATABASE_PASSWORD", "postgres")
	config.Database.Name = envOrDefault("DATABASE_NAME", "")
	config.JWTSecret = envOrDefault("JWT_SECRET", "")
	config.CORS = strings.Split(envOrDefault("CORS_ALLOW_ORIGINS", "*"), ",")
	config.Admin.Name = envOrDefault("ADMIN_USERNAME", "admin")
	config.Admin.Password = envOrDefault("ADMIN_PASSWORD", "admin")
	config.Admin.Email = envOrDefault("ADMIN_EMAIL", "admin@example.com")

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT secret не может быть пустым")
	}

	if config.Database.Name == "" {
		return nil, fmt.Errorf("Имя базы данных не может быть пустым")
	}

	if config.Environment != "dev" && config.Environment != "prod" {
		return nil, fmt.Errorf("Окружение %s не найдено", config.Environment)
	}

	return &config, nil
}

func envOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		value = defaultValue
	}

	return value
}
