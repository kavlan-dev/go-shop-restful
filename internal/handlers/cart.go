package handlers

import (
	"golang-shop-restful/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartService interface {
	CreateCart(user *models.User) error
	GetCart(user_id int) (models.Cart, error)
	AddToCart(user_id, productID int) error
	ClearCart(user_id int) error
}

func (h *Handler) GetCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	cart, err := h.service.GetCart(int(userId.(float64)))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *Handler) AddToCart(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	if err := h.service.AddToCart(int(userId.(float64)), productId); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product added to cart successfully",
	})
}

func (h *Handler) ClearCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	if err := h.service.ClearCart(int(userId.(float64))); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
