package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment   string
	ServerHost    string
	ServerPort    uint
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        uint
	JWTSecret     string
	AllowOrigins  []string
	AdminName     string
	AdminPassword string
	AdminEmail    string
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")

	if err := v.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	config := Config{
		Environment:   v.GetString("env"),
		ServerHost:    v.GetString("server.host"),
		ServerPort:    v.GetUint("server.port"),
		DBHost:        v.GetString("database.host"),
		DBUser:        v.GetString("database.user"),
		DBPassword:    v.GetString("database.password"),
		DBName:        v.GetString("database.name"),
		DBPort:        v.GetUint("database.port"),
		JWTSecret:     v.GetString("jwt.secret"),
		AllowOrigins:  v.GetStringSlice("cors.allow_origins"),
		AdminName:     v.GetString("admin.username"),
		AdminPassword: v.GetString("admin.password"),
		AdminEmail:    v.GetString("admin.email"),
	}

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT secret не может быть пустым")
	}

	if config.ServerPort == 0 {
		config.ServerPort = 8080
	}

	return &config, nil
}

func GetServerAddress(config *Config) string {
	return fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
}
