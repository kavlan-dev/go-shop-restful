package service

import (
	"go-shop-restful/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type userRepository interface {
	CreateUser(newUser *model.User) error
	FindUserByUsername(username string) (*model.User, error)
	FindUserById(userId int) (*model.User, error)
	UpdateUser(userId int, updateUser *model.User) error
}

type cartCreator interface {
	CreateCart(user *model.User) error
}

type userService struct {
	storage     userRepository
	cartCreator cartCreator
}

func NewUserService(storage userRepository, cartCreator cartCreator) *userService {
	return &userService{storage: storage, cartCreator: cartCreator}
}

func (s *userService) CreateUser(newUser *model.User) error {
	if err := newUser.Validate(); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	if err := s.storage.CreateUser(newUser); err != nil {
		return err
	}

	if err := s.cartCreator.CreateCart(newUser); err != nil {
		return err
	}

	return nil
}

func (s *userService) getUserByUsername(username string) (*model.User, error) {
	return s.storage.FindUserByUsername(username)
}

func (s *userService) AuthenticateUser(username, password string) (*model.User, error) {
	user, err := s.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UserById(userId int) (*model.User, error) {
	return s.storage.FindUserById(userId)
}

func (s *userService) UpdateUser(userId int, updateUser *model.User) error {
	return s.storage.UpdateUser(userId, updateUser)
}

func (s *userService) CreateAdminIfNotExists(adminUsername, adminEmail, adminPassword string) error {
	admin, _ := s.getUserByUsername(adminUsername)
	if admin.ID != 0 {
		return nil
	}

	adminUser := &model.User{
		Username: adminUsername,
		Password: adminPassword,
		Email:    adminEmail,
		Role:     "admin",
	}

	if err := s.CreateUser(adminUser); err != nil {
		return err
	}

	return nil
}

func (s *userService) PromoteUserToAdmin(userId int) error {
	user, err := s.UserById(userId)
	if err != nil {
		return err
	}

	if user.Role == "admin" {
		return nil
	}

	user.Role = "admin"
	if err := s.UpdateUser(userId, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) DowngradeUserToCustomer(userId int) error {
	user, err := s.UserById(userId)
	if err != nil {
		return err
	}

	if user.Role == "customer" {
		return nil
	}

	user.Role = "customer"
	if err := s.UpdateUser(userId, user); err != nil {
		return err
	}

	return nil
}
