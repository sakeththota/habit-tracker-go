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
	// router.POST("/habits", h.handleCreateHabit)
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
