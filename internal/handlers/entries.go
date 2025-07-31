package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateEntry(c echo.Context) error {
	var request models.CreateEntryRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Service.CreateEntry(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Entry created!"})
}

func (h *Handler) GetEntry(c echo.Context) error {
	id := c.Param("id")

	entry, err := h.Service.GetEntryByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, entry)
}

func (h *Handler) GetCollectionEntries(c echo.Context) error {
	id := c.Param("id")

	entries, err := h.Service.GetEntriesByCollectionID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, entries)
}

func (h *Handler) DeleteEntry(c echo.Context) error {
	id := c.Param("id")

	err := h.Service.DeleteEntry(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Entry deleted!"})
}
