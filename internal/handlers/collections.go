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
	page, err := paginationParameter("page", c.QueryParam("page"), 1, 1, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	limit, err := paginationParameter("limit", c.QueryParam("limit"), 20, 1, 100)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	collections, err := h.Service.GetCollectionsPage(siteID, currentUserID(c), page, limit)
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, collections)
}

func (h *Handler) DeleteCollection(c echo.Context) error {
	id := c.Param("id")

	err := h.Service.DeleteCollection(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Collection deleted!"})
}

func (h *Handler) UpdateCollection(c echo.Context) error {
	var request models.UpdateCollectionRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	collection, err := h.Service.UpdateCollection(c.Param("id"), currentUserID(c), request)
	if err != nil {
		return serviceError(err)
	}
	return c.JSON(http.StatusOK, collection)
}
