package service

import (
	"go-shop-restful/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type userStorage interface {
	CreateUser(newUser *model.User) error
	FindUserByUsername(username string) (*model.User, error)
	FindUserById(userId int) (*model.User, error)
	UpdateUser(userId int, updateUser *model.User) error
}

func (s *service) CreateUser(newUser *model.User) error {
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

	if err := s.CreateCart(newUser); err != nil {
		return err
	}

	return nil
}

func (s *service) getUserByUsername(username string) (*model.User, error) {
	return s.storage.FindUserByUsername(username)
}

func (s *service) AuthenticateUser(username, password string) (*model.User, error) {
	user, err := s.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) UserById(userId int) (*model.User, error) {
	return s.storage.FindUserById(userId)
}

func (s *service) UpdateUser(userId int, updateUser *model.User) error {
	return s.storage.UpdateUser(userId, updateUser)
}

func (s *service) CreateAdminIfNotExists(adminUsername, adminEmail, adminPassword string) error {
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

func (s *service) PromoteUserToAdmin(userId int) error {
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

func (s *service) DowngradeUserToCustomer(userId int) error {
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
