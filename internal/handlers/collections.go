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

	err := h.Service.CreateCollection(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Collection created!"})
}

func (h *Handler) GetCollection(c echo.Context) error {
	id := c.Param("id")

	collection, err := h.Service.GetCollectionByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, collection)
}

func (h *Handler) GetCollectionsBySiteID(c echo.Context) error {
	siteID := c.Param("id")

	collections, err := h.Service.GetCollectionsBySiteID(siteID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"collections": collections})
}

func (h *Handler) DeleteCollection(c echo.Context) error {
	id := c.Param("id")

	err := h.Service.DeleteCollection(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Collection deleted!"})
}
