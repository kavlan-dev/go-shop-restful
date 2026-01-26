package handlers

import (
	"go-shop-restful/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productService interface {
	Products(limit, offset int) (*[]models.Product, error)
	CreateProduct(product *models.Product) error
	ProductById(id int) (*models.Product, error)
	ProductByTitle(title string) (*models.Product, error)
	UpdateProduct(id int, updateProduct *models.Product) error
	DeleteProduct(id int) error
}

// TODO Добавить фильтрацию и сортировку
func (h *Handler) Products(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	var err error
	limit := 100
	offset := 0
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			h.log.Errorf("Ошибка парсинга параметра limit: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "не верно введен limit",
			})
			return
		}
		if limit > 1000 {
			limit = 1000
		}
	}
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			h.log.Errorf("Ошибка парсинга параметра offset: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "не верно введен offset",
			})
			return
		}
	}

	products, err := h.service.Products(limit, offset)
	if err != nil {
		h.log.Errorf("Ошибка при получении всех товаров: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось получить товары",
		})
		return
	}

	h.log.Debugf("Получено %d товаров", len(*products))
	c.JSON(http.StatusOK, products)
}

func (h *Handler) PostProduct(c *gin.Context) {
	var req models.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка в теле создания товара: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верное тело запроса",
		})
		return
	}

	newProduct := &models.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
	}

	if err := newProduct.Validate(); err != nil {
		h.log.Error("Некорректные данные для создания товара:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := h.service.CreateProduct(newProduct); err != nil {
		h.log.Errorf("Ошибка создания товара %s: %v", newProduct.Title, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось создать товар",
		})
		return
	}

	h.log.Debugf("Создан товар #%d с названием %s", newProduct.ID, newProduct.Title)
	c.JSON(http.StatusCreated, newProduct)
}

func (h *Handler) ProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Errorf("Ошибка парсинга ID товара: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верно введен id",
		})
		return
	}

	product, err := h.service.ProductById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Debugf("Товар #%d не найден", id)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "товар не найден",
			})
		} else {
			h.log.Errorf("Ошибка при получении товара #%d: %v", id, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "не удалось найти товар",
			})
		}
		return
	}

	h.log.Debugf("Получен товар #%d с названием %s", id, product.Title)
	c.JSON(http.StatusOK, product)
}

func (h *Handler) ProductByTitle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		h.log.Error("Пустой заголовок товара в запросе")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "отсутствует заголовок",
		})
		return
	}

	product, err := h.service.ProductByTitle(title)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Debugf("Товар с названием \"%s\" не найден", title)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "товар не найден",
			})
		} else {
			h.log.Errorf("Ошибка при получении товара \"%s\": %v", title, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "не удалось найти товар",
			})
		}
		return
	}

	h.log.Debugf("Получен товар #%d с названием %s", product.ID, title)
	c.JSON(http.StatusOK, product)
}

func (h *Handler) PutProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Errorf("Ошибка парсинга ID товара: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верно введен id",
		})
		return
	}

	var req models.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка в теле запроса для обновления товара #%d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верное тело запроса",
		})
		return
	}

	updateProduct := &models.Product{
		Category:    req.Category,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := updateProduct.Validate(); err != nil {
		h.log.Errorf("Некорректные данные для обновления товара #%d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := h.service.UpdateProduct(id, updateProduct); err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Debugf("Товар #%d не найден для обновления", id)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "товар не найден",
			})
		} else {
			h.log.Errorf("Ошибка при изменении товара #%d: %v", id, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "не удалось обновить товар",
			})
		}
		return
	}

	h.log.Debugf("Товар #%d успешно изменен", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "продукт успешно изменен",
	})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Errorf("Ошибка парсинга ID товара: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верно введен id",
		})
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			h.log.Debugf("Товар #%d не найден для удаления", id)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "товар не найден",
			})
		} else {
			h.log.Errorf("Ошибка при удалении товара #%d: %v", id, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "не удалось удалить товар",
			})
		}
		return
	}

	h.log.Debugf("Товар #%d успешно удален", id)
	c.JSON(http.StatusNoContent, nil)
}
