package postgres

import "go-shop-restful/internal/models"

func (s *storage) CreateCart(cart *models.Cart) error {
	return s.db.Create(&cart).Error
}

func (s *storage) FindCart(user_id int) (*models.Cart, error) {
	var cart models.Cart
	err := s.db.Preload("Items").Where("user_id = ?", user_id).First(&cart).Error
	return &cart, err
}

func (s *storage) FindCartItems(cart_id int) (*[]models.CartItem, error) {
	var cartItems []models.CartItem
	err := s.db.Where("cart_id = ?", cart_id).Find(&cartItems).Error
	return &cartItems, err
}

func (s *storage) ClearCart(cartItems *[]models.CartItem) error {
	return s.db.Delete(&cartItems).Error
}

func (s *storage) FindCartItem(cartId, productId int) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := s.db.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&cartItem).Error
	return &cartItem, err
}

func (s *storage) UpdateCartItem(cartItemId int, updateCartItem *models.CartItem) error {
	var cartItem models.CartItem
	if err := s.db.First(&cartItem, cartItemId).Error; err != nil {
		return err
	}
	return s.db.Model(&cartItem).Updates(&updateCartItem).Error
}

func (s *storage) CreateCartItem(cartItem *models.CartItem) error {
	return s.db.Create(&cartItem).Error
}
