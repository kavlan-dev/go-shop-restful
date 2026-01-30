package handler

import (
	"go-shop-restful/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type cartService interface {
	Cart(user_id int) (*model.Cart, error)
	AddToCart(user_id, productID int) error
	ClearCart(user_id int) error
}

func (h *handler) Cart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "не авторизован",
		})
		return
	}

	userIdUint, ok := userId.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "внутренняя ошибка сервера",
		})
		return
	}

	cart, err := h.service.Cart(int(userIdUint))
	if err != nil {
		h.log.Error("Ошибка при выводе корзины:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось получить корзину",
			"details": err,
		})
		return
	}

	h.log.Debugf("Получена корзина пользователя #%d: %v", userId, cart)
	c.JSON(http.StatusOK, cart)
}

func (h *handler) AddToCart(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "не авторизован",
		})
		return
	}

	userIdUint, ok := userId.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "внутренняя ошибка сервера",
		})
		return
	}

	if err := h.service.AddToCart(int(userIdUint), productId); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "товар не найден",
			})
		} else {
			h.log.Errorf("Ошибка при добавлении товара #%d в корзину: %v", productId, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "не удалось добавить товар в корзину",
				"details": err,
			})
		}
		return
	}

	h.log.Debugf("Товар #%d добавлен в корзину", productId)
	c.JSON(http.StatusOK, gin.H{
		"message": "товар успешно добавлен в корзину",
	})
}

func (h *handler) ClearCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "не авторизован",
		})
		return
	}

	userIdUint, ok := userId.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "внутренняя ошибка сервера",
		})
		return
	}

	if err := h.service.ClearCart(int(userIdUint)); err != nil {
		h.log.Error("Ошибка при отчистке корзины:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось отчистить корзину",
			"details": err,
		})
		return
	}

	h.log.Debug("Корзина отчищена")
	c.JSON(http.StatusNoContent, nil)
}
