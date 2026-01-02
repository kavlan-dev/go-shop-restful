package services

import (
	"golang-shop-restful/internal/models"
	"golang-shop-restful/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Services struct {
	db *gorm.DB
}

func NewServices(db *gorm.DB) *Services {
	return &Services{db: db}
}

func (s *Services) GetProducts(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Services) CreateProduct(product *models.Product) error {
	if err := s.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (s *Services) GetProductById(id int) (models.Product, error) {
	var product models.Product

	if err := s.db.First(&product, id).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *Services) UpdateProduct(id int, updateProduct *models.Product) error {
	var existingProduct models.Product
	if err := s.db.First(&existingProduct, id).Error; err != nil {
		return err
	}

	existingProduct = models.Product{
		Title:       updateProduct.Title,
		Description: updateProduct.Description,
		Price:       updateProduct.Price,
		Category:    existingProduct.Category,
		Stock:       existingProduct.Stock,
	}

	if err := s.db.Save(&existingProduct).Error; err != nil {
		return err
	}

	return nil
}

func (s *Services) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (s *Services) getUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Services) AuthenticateUser(username, password string) (models.User, error) {
	user, err := s.getUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, err
	}

	if err := s.CreateCart(&user); err != nil {
		return models.User{}, err
	}

	utils.Logger.Debug(user)
	return user, nil
}

func (s *Services) CreateCart(user *models.User) error {
	if user.Cart.UserID == 0 {
		user.Cart = models.Cart{UserID: user.ID}
		if err := s.db.Save(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *Services) GetCart(user_id int) (models.Cart, error) {
	var cart models.Cart
	if err := s.db.Preload("Items").Where("user_id = ?", user_id).First(&cart).Error; err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (s *Services) AddToCart(user_id, productID int) error {
	var user models.User
	if err := s.db.Preload("Cart").First(&user, user_id).Error; err != nil {
		return err
	}

	cart := user.Cart
	if cart.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	var product models.Product
	if err := s.db.First(&product, productID).Error; err != nil {
		return err
	}

	if product.Stock <= 0 {
		return gorm.ErrRecordNotFound
	}

	var existingCartItem models.CartItem
	if err := s.db.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&existingCartItem).Error; err == nil {
		existingCartItem.Quantity += 1
		existingCartItem.Price = product.Price * float64(existingCartItem.Quantity)
		if err := s.db.Save(&existingCartItem).Error; err != nil {
			return err
		}
	} else {
		newCartItem := models.CartItem{
			CartID:    cart.ID,
			ProductID: uint(productID),
			Quantity:  1,
			Price:     product.Price,
		}
		if err := s.db.Create(&newCartItem).Error; err != nil {
			return err
		}
	}
	product.Stock -= 1

	return nil
}
