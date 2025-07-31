package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUser(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	user, err := h.Service.GetUser(userID.(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userID := c.Param("userId")
	currentUserID := c.Get("user_id")

	if currentUserID == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	// Users can only delete their own account
	if userID != currentUserID.(string) {
		return echo.NewHTTPError(http.StatusForbidden, "Cannot delete other users")
	}

	if err := h.Service.DeleteUser(userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
