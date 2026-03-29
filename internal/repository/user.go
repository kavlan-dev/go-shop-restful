package repository

import (
	"go-shop-restful/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (s userRepository) CreateUser(newUser *model.User) error {
	return s.db.Create(&newUser).Error
}

func (s userRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := s.db.Preload("Cart").Where("username = ?", username).First(&user).Error

	return &user, err
}

func (s userRepository) FindUserById(userId int) (*model.User, error) {
	var user model.User
	err := s.db.Preload("Cart").First(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s userRepository) UpdateUser(userId int, updateUser *model.User) error {
	var user model.User
	if err := s.db.Preload("Cart").First(&user, userId).Error; err != nil {
		return err
	}

	return s.db.Model(&user).Updates(updateUser).Error
}
