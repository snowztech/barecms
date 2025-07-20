package handlers

import (
	"barecms/internal/models"
	"barecms/internal/utils"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var request models.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, h.Config.JWTSecret)
	if err != nil {
		slog.Error("Failed to generate JWT token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var request models.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Failed to bind registration request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Attempting to register user", "email", request.Email, "username", request.Username)

	// Validate JWT secret before proceeding
	if h.Config.JWTSecret == "" {
		slog.Error("JWT secret is empty")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server configuration error"})
		return
	}

	// Register the user
	if err := h.Service.Register(request); err != nil {
		slog.Error("Failed to register user", "error", err, "email", request.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Info("User registered successfully", "email", request.Email)

	// After successful registration, log the user in
	user, err := h.Service.Login(request.Email, request.Password)
	if err != nil {
		slog.Error("Failed to login after registration", "error", err, "email", request.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration successful but login failed"})
		return
	}

	slog.Info("User logged in after registration", "user_id", user.ID, "email", user.Email)

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, h.Config.JWTSecret)
	if err != nil {
		slog.Error("Failed to generate JWT token after registration", "error", err, "user_id", user.ID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	slog.Info("JWT token generated successfully", "user_id", user.ID)

	c.JSON(http.StatusCreated, gin.H{
		"token":   token,
		"user":    user,
		"message": "User created successfully",
	})
}

func (h *Handler) Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if err := h.Service.Logout(userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}
