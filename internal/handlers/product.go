package handlers

import (
	"golang-shop-restful/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService interface {
	GetProducts(limit, offset int) ([]models.Product, error)
	CreateProduct(product *models.Product) error
	GetProductById(id int) (models.Product, error)
	UpdateProduct(id int, updateProduct *models.Product) error
	DeleteProduct(id int) error
}

func (h *Handler) GetProducts(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	var err error
	limit := 100
	offset := 0
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if limit > 1000 {
			limit = 1000
		}
	}
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	products, err := h.service.GetProducts(limit, offset)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) PostProduct(c *gin.Context) {
	var req models.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newProduct := models.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
	}

	if err := h.service.CreateProduct(&newProduct); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (h *Handler) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product, err := h.service.GetProductById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) PutProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req models.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateProduct := models.Product{
		Category:    req.Category,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := h.service.UpdateProduct(id, &updateProduct); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product successfully updated",
	})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
