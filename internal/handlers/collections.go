package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCollection(c echo.Context) error {
	var req models.CreateCollectionRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.Service.CreateCollection(req, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Collection created!"})
}

func (h *Handler) GetCollection(c echo.Context) error {
	id := c.Param("id")

	collection, err := h.Service.GetCollectionByID(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, collection)
}

func (h *Handler) GetCollectionsBySiteID(c echo.Context) error {
	siteID := c.Param("id")

	collections, err := h.Service.GetCollectionsBySiteID(siteID, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"collections": collections})
}

func (h *Handler) DeleteCollection(c echo.Context) error {
	id := c.Param("id")

	err := h.Service.DeleteCollection(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Collection deleted!"})
}
