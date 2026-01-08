package services

import (
	"gorm.io/gorm"
)

type Services struct {
	db *gorm.DB
}

func NewServices(db *gorm.DB) *Services {
	return &Services{db: db}
}
