package handlers

import (
	"barecms/configs"
	"barecms/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.Service
	Config  configs.AppConfig
}

func NewHandler(service *services.Service, config configs.AppConfig) *Handler {
	return &Handler{Service: service, Config: config}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}
