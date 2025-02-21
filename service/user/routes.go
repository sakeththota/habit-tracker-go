package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sakeththota/habit-tracker-go/config"
	"github.com/sakeththota/habit-tracker-go/service/auth"
	"github.com/sakeththota/habit-tracker-go/types"
	"github.com/sakeththota/habit-tracker-go/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
}

func (h *Handler) handleLogin(c *gin.Context) {
	var payload types.LoginUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := utils.FormatValidationErrors(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Errorf("invalid habit payload: %v", formattedErrors)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("something went wrong validating habit payload: %v", err)})
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found, invalid email or password"})
		return
	}

	if !auth.ComparePasswords(u.PasswordHash, []byte(payload.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found, invalid email or password"})
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong creating token: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully logged in", "token": token})
}

func (h *Handler) handleRegister(c *gin.Context) {
	var payload types.RegisterUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := utils.FormatValidationErrors(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Errorf("invalid habit payload: %v", formattedErrors)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("something went wrong validating habit payload: %v", err)})
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("user with email %s already exists", payload.Email)})
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong hashing password: %v", err)})
		return
	}

	err = h.store.CreateUser(types.User{
		Username:     payload.Username,
		Email:        payload.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong creating user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
