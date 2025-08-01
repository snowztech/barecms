package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetSites(c echo.Context) error {
	sites, err := h.Service.GetSites()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string][]models.Site{"sites": sites})
}

func (h *Handler) GetSite(c echo.Context) error {
	id := c.Param("id")

	site, err := h.Service.GetSite(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]models.Site{"site": site})
}

func (h *Handler) GetSiteWithCollections(c echo.Context) error {
	id := c.Param("id")

	siteWithCollections, err := h.Service.GetSiteWithCollections(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, siteWithCollections)
}

func (h *Handler) CreateSite(c echo.Context) error {
	var req models.CreateSiteRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.Service.CreateSite(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Site created!"})
}

func (h *Handler) DeleteSite(c echo.Context) error {
	id := c.Param("id")
	err := h.Service.DeleteSite(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Site deleted!"})
}

func (h *Handler) GetSiteData(c echo.Context) error {
	slug := c.Param("siteSlug")

	siteData, err := h.Service.GetSiteData(slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, siteData)
}
