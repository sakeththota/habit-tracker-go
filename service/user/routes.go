package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/service/auth"
	"github.com/sakeththota/habit-tracker-go/types"
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
	c.JSON(200, gin.H{"message": "User login endpoint successfully registered"})
}

func (h *Handler) handleRegister(c *gin.Context) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("user with email %s already exists", payload.Email)})
		return
	}

	// if not create the user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.store.CreateUser(types.User{
		Username:     payload.Username,
		Email:        payload.Email,
		PasswordHash: hashedPassword, // have to hash this
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
