package service

import (
	"go-shop-restful/internal/model"

	"gorm.io/gorm"
)

type cartRepository interface {
	CreateCart(cart *model.Cart) error
	FindCart(user_id int) (*model.Cart, error)
	FindCartItems(cart_id int) ([]model.CartItem, error)
	DeleteItem(cartItem *model.CartItem) error
	ClearCart(cartItems []model.CartItem) error
	FindCartItem(cartId, productId int) (*model.CartItem, error)
	UpdateCartItem(cartItemId int, updateCartItem *model.CartItem) error
	CreateCartItem(cartItem *model.CartItem) error
}

type cartService struct {
	repository        cartRepository
	userRepository    userRepository
	productRepository productRepository
}

func NewCartService(storage cartRepository, userRepository userRepository, productRepository productRepository) *cartService {
	return &cartService{repository: storage, userRepository: userRepository, productRepository: productRepository}
}

func (s *cartService) CreateCart(user *model.User) error {
	if user.Cart.UserID != 0 {
		return nil
	}
	cart := &model.Cart{UserID: user.ID}
	if err := s.repository.CreateCart(cart); err != nil {
		return err
	}
	return nil
}

func (s *cartService) Cart(userId int) (*model.Cart, error) {
	cart, err := s.repository.FindCart(userId)
	if err != nil {
		return nil, err
	}

	for i, v := range cart.Items {
		product, err := s.productRepository.FindProductById(int(v.ProductID))
		if err != nil {
			return nil, err
		}

		cart.Items[i].Product = *product
	}

	return cart, nil
}

func (s *cartService) AddToCart(userId, productId int) error {
	user, err := s.userRepository.FindUserById(userId)
	if err != nil {
		return err
	}

	cart := user.Cart
	if cart.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	product, err := s.productRepository.FindProductById(productId)
	if err != nil {
		return err
	}

	if product.Stock <= 0 {
		return gorm.ErrRecordNotFound
	}

	cartItem, err := s.repository.FindCartItem(int(cart.ID), productId)
	if err == nil {
		cartItem.Quantity += 1
		cartItem.Price = product.Price * float64(cartItem.Quantity)
		if err := s.repository.UpdateCartItem(int(cartItem.ID), cartItem); err != nil {
			return err
		}
	} else {
		newCartItem := &model.CartItem{
			CartID:    cart.ID,
			ProductID: uint(productId),
			Quantity:  1,
			Price:     product.Price,
		}
		if err := s.repository.CreateCartItem(newCartItem); err != nil {
			return err
		}
	}
	product.Stock -= 1
	if err := s.productRepository.UpdateProduct(productId, product); err != nil {
		return err
	}

	return nil
}

func (s *cartService) DeleteItem(user_id, itemID int) error {
	cart, err := s.repository.FindCart(user_id)
	if err != nil {
		return err
	}

	cartItem, err := s.repository.FindCartItem(int(cart.ID), itemID)
	if err != nil {
		return err
	}

	return s.repository.DeleteItem(cartItem)
}

func (s *cartService) ClearCart(user_id int) error {
	cart, err := s.repository.FindCart(user_id)
	if err != nil {
		return err
	}

	cartItems, err := s.repository.FindCartItems(int(cart.ID))
	if err != nil {
		return err
	}

	return s.repository.ClearCart(cartItems)
}
