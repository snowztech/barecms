package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHTTPErrorHandlerPreservesStructuredPayload(t *testing.T) {
	e := echo.New()
	e.HTTPErrorHandler = HTTPErrorHandler
	e.GET("/", func(echo.Context) error {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{"error": echo.Map{"code": "validation_failed"}})
	})
	response := httptest.NewRecorder()
	e.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/", nil))

	var body map[string]map[string]any
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["error"]["code"] != "validation_failed" {
		t.Fatalf("unexpected error response: %s", response.Body.String())
	}
}
