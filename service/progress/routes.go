package progress

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/service/auth"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Handler struct {
	store      types.ProgressStore
	userStore  types.UserStore
	habitStore types.HabitStore
}

func NewHandler(store types.ProgressStore, userStore types.UserStore, habitStore types.HabitStore) *Handler {
	return &Handler{store: store, userStore: userStore, habitStore: habitStore}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/progress/hello", h.handleGetProgressHello)
	router.GET("/progress/:id", auth.WithJWTAuth(h.handleGetProgress, h.userStore))
}

func (h *Handler) handleGetProgress(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)
	habitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit, err := h.habitStore.GetHabitById(habitID)
	if err != nil {
		log.Printf("failed to get habit by id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if habit.UserID != userID {
		log.Printf("failed to get progress of habit with id: %v", err)
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Errorf("permission denied")})
		return
	}

	progress, err := h.store.GetProgress(habitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

func (h *Handler) handleGetProgressHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
