package handlers

import (
	"barecms/internal/services"
	"errors"
	"mime"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadMedia(c echo.Context) error {
	siteID := c.Param("siteId")
	header, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "file form field is required")
	}
	source, err := header.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not read uploaded file")
	}
	defer func() { _ = source.Close() }()

	file, err := h.Service.UploadMedia(siteID, currentUserID(c), header.Filename, source)
	if err != nil {
		if errors.Is(err, services.ErrFileTooLarge) {
			return echo.NewHTTPError(http.StatusRequestEntityTooLarge, err.Error())
		}
		if errors.Is(err, services.ErrUnsupportedFile) {
			return echo.NewHTTPError(http.StatusUnsupportedMediaType, err.Error())
		}
		return serviceError(err)
	}
	return c.JSON(http.StatusCreated, file)
}

func (h *Handler) ListMedia(c echo.Context) error {
	page, err := paginationParameter("page", c.QueryParam("page"), 1, 1, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	limit, err := paginationParameter("limit", c.QueryParam("limit"), 50, 1, 100)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	files, err := h.Service.ListMediaPage(c.Param("siteId"), currentUserID(c), page, limit)
	if err != nil {
		return serviceError(err)
	}
	return c.JSON(http.StatusOK, files)
}

func (h *Handler) GetMedia(c echo.Context) error {
	file, path, err := h.Service.GetMedia(c.Param("fileId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "file not found")
	}
	c.Response().Header().Set(echo.HeaderContentType, file.MIMEType)
	c.Response().Header().Set(echo.HeaderContentDisposition, mime.FormatMediaType("inline", map[string]string{"filename": file.OriginalName}))
	return c.File(path)
}

func (h *Handler) DeleteMedia(c echo.Context) error {
	if err := h.Service.DeleteMedia(c.Param("fileId"), currentUserID(c)); err != nil {
		return serviceError(err)
	}
	return c.NoContent(http.StatusNoContent)
}
