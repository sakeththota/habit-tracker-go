package habit

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sakeththota/habit-tracker-go/service/auth"
	"github.com/sakeththota/habit-tracker-go/types"
	"github.com/sakeththota/habit-tracker-go/utils"
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
	router.DELETE("/habits/:id", auth.WithJWTAuth(auth.ValidateOwnership(h.handleDeleteHabit, h.store), h.userStore))
}

func (h *Handler) handleGetHabits(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)

	hs, err := h.store.GetHabits(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong getting habit: %v", err)})
		return
	}

	c.JSON(http.StatusOK, hs)
}

func (h *Handler) handleCreateHabit(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)

	var payload types.CreateHabitPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := utils.FormatValidationErrors(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Errorf("invalid habit payload: %v", formattedErrors)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("something went wrong validating habit payload: %v", err)})
		return
	}

	err := h.store.CreateHabit(types.Habit{
		UserID:      userID,
		Title:       payload.Title,
		Description: payload.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong creating habit: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *Handler) handleDeleteHabit(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)
	habitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid habit id: %v", err)})
		return
	}

	err = h.store.DeleteHabit(userID, habitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong deleting habit: %v", err)})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
