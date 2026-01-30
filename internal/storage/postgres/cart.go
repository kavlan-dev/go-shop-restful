package postgres

import "go-shop-restful/internal/model"

func (s *storage) CreateCart(cart *model.Cart) error {
	return s.db.Create(&cart).Error
}

func (s *storage) FindCart(user_id int) (*model.Cart, error) {
	var cart model.Cart
	err := s.db.Preload("Items").Where("user_id = ?", user_id).First(&cart).Error

	return &cart, err
}

func (s *storage) FindCartItems(cart_id int) (*[]model.CartItem, error) {
	var cartItems []model.CartItem
	err := s.db.Where("cart_id = ?", cart_id).Find(&cartItems).Error

	return &cartItems, err
}

func (s *storage) ClearCart(cartItems *[]model.CartItem) error {
	return s.db.Delete(&cartItems).Error
}

func (s *storage) FindCartItem(cartId, productId int) (*model.CartItem, error) {
	var cartItem model.CartItem
	err := s.db.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&cartItem).Error

	return &cartItem, err
}

func (s *storage) UpdateCartItem(cartItemId int, updateCartItem *model.CartItem) error {
	var cartItem model.CartItem
	if err := s.db.First(&cartItem, cartItemId).Error; err != nil {
		return err
	}

	return s.db.Model(&cartItem).Updates(&updateCartItem).Error
}

func (s *storage) CreateCartItem(cartItem *model.CartItem) error {
	return s.db.Create(&cartItem).Error
}
