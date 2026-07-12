package router

import (
	"barecms/configs"
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testConfig() configs.AppConfig {
	return configs.AppConfig{
		Env:                    "production",
		JWTSecret:              "a-secure-secret-with-at-least-32-characters",
		MaxRequestBody:         "16B",
		AuthRateLimitPerMinute: 1,
	}
}

func TestSecurityHeaders(t *testing.T) {
	router := Setup(nil, testConfig())
	request := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	request.TLS = &tls.ConnectionState{}
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
	if got := response.Header().Get("X-Frame-Options"); got != "DENY" {
		t.Fatalf("expected X-Frame-Options DENY, got %q", got)
	}
	if got := response.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Fatalf("expected nosniff, got %q", got)
	}
	if got := response.Header().Get("Strict-Transport-Security"); got == "" {
		t.Fatal("expected HSTS in production")
	}
}

func TestRequestBodyLimit(t *testing.T) {
	router := Setup(nil, testConfig())
	request := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(`{"email":"far-too-large@example.com"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("expected 413, got %d", response.Code)
	}
}

func TestAuthRateLimit(t *testing.T) {
	config := testConfig()
	config.MaxRequestBody = "2M"
	router := Setup(nil, config)

	first := httptest.NewRecorder()
	router.ServeHTTP(first, httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(`{`)))
	if first.Code == http.StatusTooManyRequests {
		t.Fatal("first auth request should not be rate limited")
	}

	second := httptest.NewRecorder()
	router.ServeHTTP(second, httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(`{`)))
	if second.Code != http.StatusTooManyRequests {
		t.Fatalf("expected second auth request to return 429, got %d", second.Code)
	}
}
