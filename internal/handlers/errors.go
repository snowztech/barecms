package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	if payload, ok := he.Message.(echo.Map); ok {
		_ = c.JSON(he.Code, payload)
		return
	}
	_ = c.JSON(he.Code, echo.Map{"error": echo.Map{"code": http.StatusText(he.Code), "message": he.Message}})
}
