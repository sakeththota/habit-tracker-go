package habit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Handler struct {
	store types.HabitStore
}

func NewHandler(store types.HabitStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/habits", h.handleGetHabits)
	// router.POST("/habits", h.handleCreateHabit)
}

func (h *Handler) handleGetHabits(c *gin.Context) {
	hs, err := h.store.GetHabits(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hs)
}
