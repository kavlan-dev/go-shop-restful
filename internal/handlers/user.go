package handlers

import (
	"golang-shop-restful/internal/models"
	"golang-shop-restful/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(user *models.User) error
	AuthenticateUser(username, password string) (models.User, error)
}

func (h *Handler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	role := req.Role
	if role == "" {
		role = "customer"
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     role,
	}

	if err := h.service.CreateUser(user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := h.service.CreateCart(user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user successfully created",
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.service.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		Token: token,
	})
}
