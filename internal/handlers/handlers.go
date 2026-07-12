package handlers

import (
	"barecms/configs"
	"barecms/internal/services"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func currentUserID(c echo.Context) string {
	userID, _ := c.Get("user_id").(string)
	return userID
}

func serviceError(err error) error {
	if errors.Is(err, services.ErrForbidden) {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

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
