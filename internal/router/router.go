package router

import (
	"barecms/configs"
	"barecms/internal/handlers"
	"barecms/internal/middlewares"
	"barecms/internal/services"
	"barecms/ui"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(service *services.Service, config configs.AppConfig) *echo.Echo {
	r := echo.New()

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
	api := r.Group("/api")
	api.GET("/health", h.Health)

	// Public site data endpoint
	api.GET("/:siteSlug/data", h.GetSiteData)

	// Auth routes (public)
	api.POST("/auth/register", h.Register)
	api.POST("/auth/login", h.Login)

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
	protected.DELETE("/sites/:id", h.DeleteSite)

	// Collections routes
	protected.POST("/collections", h.CreateCollection)
	protected.GET("/collections/:id", h.GetCollection)
	protected.GET("/collections/:id/entries", h.GetCollectionEntries)
	protected.GET("/collections/site/:id", h.GetCollectionsBySiteID)
	protected.DELETE("/collections/:id", h.DeleteCollection)

	// Entries routes
	protected.POST("/entries", h.CreateEntry)
	protected.GET("/entries/:id", h.GetEntry)
	protected.DELETE("/entries/:id", h.DeleteEntry)

	return r
}
