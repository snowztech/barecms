package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetSites(c echo.Context) error {
	page, err := paginationParameter("page", c.QueryParam("page"), 1, 1, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	limit, err := paginationParameter("limit", c.QueryParam("limit"), 20, 1, 100)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	sites, err := h.Service.GetSitesPage(currentUserID(c), page, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, sites)
}

func (h *Handler) GetSite(c echo.Context) error {
	id := c.Param("id")

	site, err := h.Service.GetSite(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, map[string]models.Site{"site": site})
}

func (h *Handler) GetSiteWithCollections(c echo.Context) error {
	id := c.Param("id")

	siteWithCollections, err := h.Service.GetSiteWithCollections(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}

	return c.JSON(http.StatusOK, siteWithCollections)
}

func (h *Handler) CreateSite(c echo.Context) error {
	var req models.CreateSiteRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.Service.CreateSite(req, currentUserID(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Site created!"})
}

func (h *Handler) DeleteSite(c echo.Context) error {
	id := c.Param("id")
	err := h.Service.DeleteSite(id, currentUserID(c))
	if err != nil {
		return serviceError(err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Site deleted!"})
}

func (h *Handler) UpdateSite(c echo.Context) error {
	var request models.UpdateSiteRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	site, err := h.Service.UpdateSite(c.Param("id"), currentUserID(c), request)
	if err != nil {
		return serviceError(err)
	}
	return c.JSON(http.StatusOK, map[string]models.Site{"site": site})
}

func (h *Handler) GetSiteData(c echo.Context) error {
	slug := c.Param("siteSlug")

	siteData, err := h.Service.GetSiteData(slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, siteData)
}
