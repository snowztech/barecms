package middlewares

import (
	"barecms/configs"
	"barecms/internal/models"
	"barecms/internal/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SiteAuthMiddleware checks if user has access to a specific site
func SiteAuthMiddleware(db *gorm.DB, config configs.AppConfig, requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// First check basic auth
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header required"})
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid authorization format"})
			}

			token := tokenParts[1]
			claims, err := utils.ValidateJWT(token, config.JWTSecret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
			}

			// Get site ID from URL params
			siteID := c.Param("siteId")
			if siteID == "" {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": "Site ID required"})
			}

			// Check site permissions
			hasPermission, userRole, err := checkSitePermission(db, claims.UserID, siteID, requiredRole)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to check permissions"})
			}

			if !hasPermission {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "Insufficient permissions"})
			}

			// Set context values
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("site_id", siteID)
			c.Set("user_role", userRole)

			return next(c)
		}
	}
}

// checkSitePermission verifies if user has required role for site
func checkSitePermission(db *gorm.DB, userID, siteID, requiredRole string) (bool, string, error) {
	// First check if user is site owner
	var site models.Site
	err := db.Where("id = ? AND owner_id = ?", siteID, userID).First(&site).Error
	if err == nil {
		return true, models.RoleOwner, nil // Owner has all permissions
	}

	// Check site_users table for collaborator access
	var siteUser models.SiteUser
	err = db.Where("site_id = ? AND user_id = ? AND joined_at IS NOT NULL", siteID, userID).First(&siteUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, "", nil // No access
		}
		return false, "", err // Database error
	}

	// Check if user's role meets requirements
	hasPermission := roleHasPermission(siteUser.Role, requiredRole)
	return hasPermission, siteUser.Role, nil
}

// roleHasPermission checks if user role meets required permission level
func roleHasPermission(userRole, requiredRole string) bool {
	roleHierarchy := map[string]int{
		models.RoleViewer: 1,
		models.RoleEditor: 2,
		models.RoleOwner:  3,
	}

	userLevel, userExists := roleHierarchy[userRole]
	requiredLevel, requiredExists := roleHierarchy[requiredRole]

	if !userExists || !requiredExists {
		return false
	}

	return userLevel >= requiredLevel
}

// Helper middleware functions for common permission levels
func SiteViewerAuth(db *gorm.DB, config configs.AppConfig) echo.MiddlewareFunc {
	return SiteAuthMiddleware(db, config, models.RoleViewer)
}

func SiteEditorAuth(db *gorm.DB, config configs.AppConfig) echo.MiddlewareFunc {
	return SiteAuthMiddleware(db, config, models.RoleEditor)
}

func SiteOwnerAuth(db *gorm.DB, config configs.AppConfig) echo.MiddlewareFunc {
	return SiteAuthMiddleware(db, config, models.RoleOwner)
}