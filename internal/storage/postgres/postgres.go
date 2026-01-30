package postgres

import (
	"fmt"
	"go-shop-restful/internal/config"
	"go-shop-restful/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *storage {
	return &storage{db: db}
}

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Product{}, &model.User{}, &model.Cart{}, &model.CartItem{}); err != nil {
		return nil, err
	}

	return db, nil
}
