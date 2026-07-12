package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// InviteUserToSite invites a user to collaborate on a site
func (h *Handler) InviteUserToSite(c echo.Context) error {
	siteID := c.Param("siteId")
	userID := c.Get("user_id").(string)

	var request models.InviteUserRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Basic validation
	if request.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email is required")
	}
	if request.Role != models.RoleEditor && request.Role != models.RoleViewer {
		return echo.NewHTTPError(http.StatusBadRequest, "Role must be 'editor' or 'viewer'")
	}

	invitation, err := h.SiteUserService.InviteUser(siteID, userID, request.Email, request.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":    "User invited successfully",
		"invitation": invitation,
	})
}

// GetSiteUsers returns all users for a site
func (h *Handler) GetSiteUsers(c echo.Context) error {
	siteID := c.Param("siteId")

	users, err := h.SiteUserService.GetSiteUsers(siteID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch site users")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
	})
}

// RemoveUserFromSite removes a user from a site
func (h *Handler) RemoveUserFromSite(c echo.Context) error {
	siteID := c.Param("siteId")
	userToRemoveID := c.Param("userId")
	ownerID := c.Get("user_id").(string)

	err := h.SiteUserService.RemoveUser(siteID, ownerID, userToRemoveID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User removed successfully",
	})
}

// UpdateUserRole updates a user's role in a site
func (h *Handler) UpdateUserRole(c echo.Context) error {
	siteID := c.Param("siteId")
	userID := c.Param("userId")
	ownerID := c.Get("user_id").(string)

	var request struct {
		Role string `json:"role"`
	}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	err := h.SiteUserService.UpdateUserRole(siteID, ownerID, userID, request.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User role updated successfully",
	})
}