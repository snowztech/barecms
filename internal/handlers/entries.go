package handlers

import (
	"barecms/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateEntry(c echo.Context) error {
	var request models.CreateEntryRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Service.CreateEntry(&request, currentUserID(c)); err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Entry created!"})
}

func (h *Handler) GetEntry(c echo.Context) error {
	id := c.Param("id")

	entry, err := h.Service.GetEntryByID(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, entry)
}

func (h *Handler) GetCollectionEntries(c echo.Context) error {
	id := c.Param("id")
	page, err := paginationParameter("page", c.QueryParam("page"), 1, 1, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	limit, err := paginationParameter("limit", c.QueryParam("limit"), 20, 1, 100)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	entries, err := h.Service.GetEntriesPage(id, currentUserID(c), page, limit)
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, entries)
}

func paginationParameter(name, raw string, fallback, minimum, maximum int) (int, error) {
	if raw == "" {
		return fallback, nil
	}
	value, err := strconv.Atoi(raw)
	if err != nil || value < minimum || (maximum > 0 && value > maximum) {
		if maximum == 0 {
			return 0, fmt.Errorf("%s must be an integer of at least %d", name, minimum)
		}
		return 0, fmt.Errorf("%s must be an integer between %d and %d", name, minimum, maximum)
	}
	return value, nil
}

func (h *Handler) DeleteEntry(c echo.Context) error {
	id := c.Param("id")

	err := h.Service.DeleteEntry(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Entry deleted!"})
}

func (h *Handler) UpdateEntry(c echo.Context) error {
	var request models.UpdateEntryRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	entry, err := h.Service.UpdateEntry(c.Param("id"), currentUserID(c), request)
	if err != nil {
		return serviceError(err)
	}
	return c.JSON(http.StatusOK, entry)
}
