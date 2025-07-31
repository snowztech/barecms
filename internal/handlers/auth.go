package handlers

import (
	"barecms/internal/models"
	"barecms/internal/utils"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	var request models.LoginRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.Service.Login(request.Email, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, h.Config.JWTSecret)
	if err != nil {
		slog.Error("Failed to generate JWT token", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"token": token,
		"user":  user,
	})
}

func (h *Handler) Register(c echo.Context) error {
	var request models.RegisterRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	slog.Info("Attempting to register user", "email", request.Email, "username", request.Username)

	// Validate JWT secret before proceeding
	if h.Config.JWTSecret == "" {
		slog.Error("JWT secret is empty")
		return echo.NewHTTPError(http.StatusInternalServerError, "Server configuration error")
	}

	// Register the user
	if err := h.Service.Register(request); err != nil {
		slog.Error("Failed to register user", "error", err, "email", request.Email)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	slog.Info("User registered successfully", "email", request.Email)

	// After successful registration, log the user in
	user, err := h.Service.Login(request.Email, request.Password)
	if err != nil {
		slog.Error("Failed to login after registration", "error", err, "email", request.Email)
		return echo.NewHTTPError(http.StatusInternalServerError, "Registration successful but login failed")
	}

	slog.Info("User logged in after registration", "user_id", user.ID, "email", user.Email)

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, h.Config.JWTSecret)
	if err != nil {
		slog.Error("Failed to generate JWT token after registration", "error", err, "user_id", user.ID)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	slog.Info("JWT token generated successfully", "user_id", user.ID)

	return c.JSON(http.StatusCreated, map[string]any{
		"token":   token,
		"user":    user,
		"message": "User created successfully",
	})
}

func (h *Handler) Logout(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	if err := h.Service.Logout(userID.(string)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User logged out successfully",
	})
}
