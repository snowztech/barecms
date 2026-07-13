package router

import (
	"barecms/configs"
	"barecms/internal/handlers"
	"barecms/internal/middlewares"
	"barecms/internal/services"
	"barecms/ui"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func Setup(service *services.Service, config configs.AppConfig) *echo.Echo {
	r := echo.New()
	r.HTTPErrorHandler = handlers.HTTPErrorHandler
	r.Use(middleware.BodyLimit(config.MaxRequestBody))
	r.Use(middleware.SecureWithConfig(securityHeaders(config)))

	if config.Env == "dev" {
		r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:5173", "http://localhost:5172"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		}))
	}

	// Serve the built frontend files
	r.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: ui.BuildHTTPFS(),
		HTML5:      true,
	}))

	h := handlers.NewHandler(service, config)
	r.GET("/healthz", h.Health)
	r.GET("/readyz", h.Readiness)
	api := r.Group("/api")
	api.GET("/health", h.Health)

	// Public site data endpoint
	api.GET("/:siteSlug/data", h.GetSiteData)
	api.GET("/files/:fileId", h.GetMedia)

	// Auth routes (public)
	authRateLimiter := middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      rate.Limit(float64(config.AuthRateLimitPerMinute) / 60),
			Burst:     config.AuthRateLimitPerMinute,
			ExpiresIn: 5 * time.Minute,
		},
	))
	api.POST("/auth/register", h.Register, authRateLimiter)
	api.POST("/auth/login", h.Login, authRateLimiter)

	// Protected routes
	protected := api.Group("")
	protected.Use(middlewares.AuthMiddleware(config))

	// User Management
	protected.GET("/user", h.GetUser)
	protected.DELETE("/user/:userId", h.DeleteUser)

	// Sites routes
	protected.GET("/sites", h.GetSites)
	protected.GET("/sites/:id", h.GetSite)
	protected.GET("/sites/:id/collections", h.GetSiteWithCollections)
	protected.POST("/sites", h.CreateSite)
	protected.PUT("/sites/:id", h.UpdateSite)
	protected.DELETE("/sites/:id", h.DeleteSite)
	protected.GET("/sites/:siteId/files", h.ListMedia)
	protected.POST("/sites/:siteId/files", h.UploadMedia)
	protected.DELETE("/files/:fileId", h.DeleteMedia)

	// Collections routes
	protected.POST("/collections", h.CreateCollection)
	protected.GET("/collections/:id", h.GetCollection)
	protected.PUT("/collections/:id", h.UpdateCollection)
	protected.GET("/collections/:id/entries", h.GetCollectionEntries)
	protected.GET("/collections/site/:id", h.GetCollectionsBySiteID)
	protected.DELETE("/collections/:id", h.DeleteCollection)

	// Entries routes
	protected.POST("/entries", h.CreateEntry)
	protected.GET("/entries/:id", h.GetEntry)
	protected.PUT("/entries/:id", h.UpdateEntry)
	protected.DELETE("/entries/:id", h.DeleteEntry)

	return r
}

func securityHeaders(config configs.AppConfig) middleware.SecureConfig {
	secure := middleware.DefaultSecureConfig
	secure.XFrameOptions = "DENY"
	secure.ContentTypeNosniff = "nosniff"
	secure.ReferrerPolicy = "strict-origin-when-cross-origin"
	if config.Env == "prod" || config.Env == "production" {
		secure.HSTSMaxAge = 31536000
	}
	return secure
}
