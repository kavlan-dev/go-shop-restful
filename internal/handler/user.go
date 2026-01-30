package handler

import (
	"go-shop-restful/internal/model"
	"go-shop-restful/internal/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userService interface {
	CreateCart(user *model.User) error
	CreateUser(user *model.User) error
	AuthenticateUser(username, password string) (*model.User, error)
	PromoteUserToAdmin(userID int) error
	DowngradeUserToCustomer(userID int) error
}

func (h *handler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка в теле запроса регистрации: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верное тело запроса",
		})
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := h.service.CreateUser(user); err != nil {
		h.log.Errorf("Ошибка при создании пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось создать пользователя",
		})
		return
	}

	h.log.Debugf("Успешное создание пользователя #%d", user.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "пользователь успешно создан",
	})
}

func (h *handler) Login(c *gin.Context) {
	var req model.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка в теле запроса логина: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верное тело запроса",
		})
		return
	}

	user, err := h.service.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		h.log.Errorf("Ошибка авторизации пользователя %s: %v", req.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "неверное имя пользователя или пароль",
		})
		return
	}

	token, err := util.GenerateJWT(user.ID, user.Role)
	if err != nil {
		h.log.Errorf("Ошибка генерации токена для пользователя #%d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось сгенерировать JWT токен",
		})
		return
	}

	h.log.Debugf("Пользователь #%d успешно вошел", user.ID)
	c.JSON(http.StatusOK, model.AuthResponse{
		Token: token,
	})
}

func (h *handler) PromoteToAdmin(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Errorf("Ошибка парсинга ID пользователя: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верно введен id",
		})
		return
	}

	if err := h.service.PromoteUserToAdmin(userID); err != nil {
		h.log.Errorf("Ошибка повышения пользователя #%d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось назначить пользователя администратором",
		})
		return
	}

	h.log.Debugf("Пользователь #%d повышен до администратора", userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "пользователь повышен до администратора",
	})
}

func (h *handler) DowngradeToCustomer(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Errorf("Ошибка парсинга ID пользователя: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верно введен id",
		})
		return
	}

	if err := h.service.DowngradeUserToCustomer(userID); err != nil {
		h.log.Errorf("Ошибка понижения пользователя #%d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось понизить пользователя",
		})
		return
	}

	h.log.Debugf("Пользователь #%d понижен до обычного пользователя", userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "пользователь понижен до обычного пользователя",
	})
}
