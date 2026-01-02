package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string  `json:"title" gorm:"size:100;not null" binding:"required,min=1,max=100"`
	Description string  `json:"description" gorm:"size:1000" binding:"max=1000"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"required,min=0"`
	Category    string  `json:"category" gorm:"size:50" binding:"max=50"`
	Stock       uint    `json:"stock" gorm:"not null;default:0" binding:"min=0"`
}
