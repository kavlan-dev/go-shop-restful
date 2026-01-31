package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Server      serverConfig
	Database    databaseConfig
	JWTSecret   string
	CORS        []string
	Admin       adminConfig
}

type serverConfig struct {
	Host string
	Port uint
}

type databaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     uint
}

type adminConfig struct {
	Name     string
	Password string
	Email    string
}

func InitConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")

	if err := v.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	config := Config{
		Environment: v.GetString("env"),
		Server: serverConfig{
			Host: v.GetString("server.host"),
			Port: v.GetUint("server.port"),
		},
		Database: databaseConfig{
			Host:     v.GetString("database.host"),
			Port:     v.GetUint("database.port"),
			User:     v.GetString("database.user"),
			Password: v.GetString("database.password"),
			Name:     v.GetString("database.name"),
		},
		JWTSecret: v.GetString("jwt.secret"),
		CORS:      v.GetStringSlice("cors.allow_origins"),
		Admin: adminConfig{
			Name:     v.GetString("admin.username"),
			Password: v.GetString("admin.password"),
			Email:    v.GetString("admin.email"),
		},
	}

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT secret не может быть пустым")
	}

	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}

	if config.Environment != "dev" && config.Environment != "prod" {
		return nil, fmt.Errorf("Окружение %s не найдено", config.Environment)
	}

	return &config, nil
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
