package services

import (
	"go-shop-restful/internal/models"

	"gorm.io/gorm"
)

type cartStorage interface {
	CreateCart(cart *models.Cart) error
	FindCart(user_id int) (*models.Cart, error)
	FindCartItems(cart_id int) (*[]models.CartItem, error)
	ClearCart(cartItems *[]models.CartItem) error
	FindCartItem(cartId, productId int) (*models.CartItem, error)
	UpdateCartItem(cartItemId int, updateCartItem *models.CartItem) error
	CreateCartItem(cartItem *models.CartItem) error
}

func (s *service) CreateCart(user *models.User) error {
	if user.Cart.UserID != 0 {
		return nil
	}
	cart := models.Cart{UserID: user.ID}
	if err := s.storage.CreateCart(&cart); err != nil {
		return err
	}
	return nil
}

func (s *service) Cart(userId int) (*models.Cart, error) {
	return s.storage.FindCart(userId)
}

func (s *service) AddToCart(userId, productId int) error {
	user, err := s.GetUserById(userId)
	if err != nil {
		return err
	}

	cart := user.Cart
	if cart.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	product, err := s.ProductById(productId)
	if err != nil {
		return err
	}

	if product.Stock <= 0 {
		return gorm.ErrRecordNotFound
	}

	cartItem, err := s.storage.FindCartItem(int(cart.ID), productId)
	if err == nil {
		cartItem.Quantity += 1
		cartItem.Price = product.Price * float64(cartItem.Quantity)
		if err := s.storage.UpdateCartItem(int(cartItem.ID), cartItem); err != nil {
			return err
		}
	} else {
		newCartItem := models.CartItem{
			CartID:    cart.ID,
			ProductID: uint(productId),
			Quantity:  1,
			Price:     product.Price,
		}
		if err := s.storage.CreateCartItem(&newCartItem); err != nil {
			return err
		}
	}
	product.Stock -= 1
	if err := s.UpdateProduct(productId, product); err != nil {
		return err
	}

	return nil
}

func (s *service) ClearCart(user_id int) error {
	cart, err := s.storage.FindCart(user_id)
	if err != nil {
		return err
	}

	cartItems, err := s.storage.FindCartItems(int(cart.ID))
	if err != nil {
		return err
	}

	return s.storage.ClearCart(cartItems)
}
