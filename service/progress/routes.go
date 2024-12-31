package progress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Handler struct {
	store     types.ProgressStore
	userStore types.UserStore
}

func NewHandler(store types.ProgressStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/progress/hello", h.handleGetProgressHello)
}

func (h *Handler) handleGetProgressHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
