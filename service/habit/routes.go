package habit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/service/auth"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Handler struct {
	store     types.HabitStore
	userStore types.UserStore
}

func NewHandler(store types.HabitStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/habits", auth.WithJWTAuth(h.handleGetHabits, h.userStore))
	router.POST("/habits", auth.WithJWTAuth(h.handleCreateHabit, h.userStore))
}

func (h *Handler) handleGetHabits(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)

	hs, err := h.store.GetHabits(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hs)
}

func (h *Handler) handleCreateHabit(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)

	var payload types.CreateHabitPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.store.CreateHabit(types.Habit{
		UserID:      userID,
		Title:       payload.Title,
		Description: payload.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
