package model

import "gorm.io/gorm"

const (
	OrderStatusPending    = "pending"
	OrderStatusProcessing = "processing"
	OrderStatusShipped    = "shipped"
	OrderStatusDelivered  = "delivered"
	OrderStatusCancelled  = "cancelled"
)

type Cart struct {
	gorm.Model
	UserID uint       `json:"user_id" gorm:"unique;not null"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null" binding:"min=1"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"min=0"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}

// TODO Добавить заказы
// type Order struct {
// 	gorm.Model
// 	UserID uint        `json:"user_id" gorm:"not null"`
// 	Total  float64     `json:"total" gorm:"type:decimal(10,2);not null"`
// 	Status string      `json:"status" gorm:"size:20;default:pending" binding:"oneof=pending processing shipped delivered cancelled"`
// 	Items  []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
// }

// type OrderItem struct {
// 	gorm.Model
// 	OrderID   uint    `json:"order_id" gorm:"not null"`
// 	ProductID uint    `json:"product_id" gorm:"not null"`
// 	Quantity  int     `json:"quantity" gorm:"not null" binding:"min=1"`
// 	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"min=0"`
// 	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
// }
