package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const publicContentCacheControl = "public, max-age=60, stale-while-revalidate=300"

func (h *Handler) GetPublicEntries(c echo.Context) error {
	page, err := paginationParameter("page", c.QueryParam("page"), 1, 1, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	limit, err := paginationParameter("limit", c.QueryParam("limit"), 20, 1, 100)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	entries, err := h.Service.GetPublicEntries(c.Param("siteSlug"), c.Param("collectionSlug"), page, limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "content not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "could not load content")
	}
	c.Response().Header().Set(echo.HeaderCacheControl, publicContentCacheControl)
	return c.JSON(http.StatusOK, entries)
}

func (h *Handler) GetPublicEntry(c echo.Context) error {
	entry, err := h.Service.GetPublicEntry(c.Param("siteSlug"), c.Param("collectionSlug"), c.Param("entryId"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "content not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "could not load content")
	}
	c.Response().Header().Set(echo.HeaderCacheControl, publicContentCacheControl)
	return c.JSON(http.StatusOK, entry)
}
