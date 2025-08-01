package handlers

import (
	"barecms/configs"
	"barecms/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *services.Service
	Config  configs.AppConfig
}

func NewHandler(service *services.Service, config configs.AppConfig) *Handler {
	return &Handler{Service: service, Config: config}
}

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "up"})
}
