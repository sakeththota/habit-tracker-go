package progress

import (
	"fmt"
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
	router.GET("/progress/:id", auth.ValidateOwnership(auth.WithJWTAuth(h.handleGetProgress, h.userStore), h.habitStore))
	router.POST("/progress/:id/:date", auth.ValidateOwnership(auth.WithJWTAuth(h.handleMarkComplete, h.userStore), h.habitStore))
}

func (h *Handler) handleMarkComplete(c *gin.Context) {
	habitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid habit id: %v", err)})
		return
	}
	date := c.Param("date")

	err = h.store.CreateCompletion(habitID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong marking habit complete: %v", err)})
	}
}

func (h *Handler) handleGetProgress(c *gin.Context) {
	habitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid habit id: %v", err)})
		return
	}

	progress, err := h.store.GetProgress(habitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong getting progress: %v", err)})
		return
	}

	c.JSON(http.StatusOK, progress)
}
