package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Handler struct {
	store types.HealthStore
}

func NewHandler(store types.HealthStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/healthz", h.handleHealthCheck)
	router.GET("/live", h.handleLive)
	router.GET("ready", h.handleReady)
}

func (h *Handler) handleHealthCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "ok"
	err := h.store.PingDB(ctx)
	if err != nil {
		dbStatus = "down"
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "degraded",
			"details": gin.H{"db": dbStatus, "error": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"details": gin.H{"db": dbStatus},
	})
}

func (h *Handler) handleLive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

func (h *Handler) handleReady(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	err := h.store.PingDB(ctx)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unavailable",
			"db":     "down",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"db":     "ok",
	})
}
