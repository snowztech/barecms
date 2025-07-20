package router

import (
	"barecms/configs"
	"barecms/internal/handlers"
	"barecms/internal/middlewares"
	"barecms/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Setup(service *services.Service, config configs.AppConfig) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if config.Env == "dev" {
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:5173", "http://localhost:5172"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		}))
	}

	// Serve static files
	router.Use(static.Serve("/", static.LocalFile("./ui/dist", true)))
	// Fallback to index.html for client-side routing
	router.NoRoute(func(c *gin.Context) {
		c.File("./ui/dist/index.html")
	})

	h := handlers.NewHandler(service, config)
	api := router.Group("/api")
	{
		api.GET("/health", h.Health)

		// Public site data endpoint
		api.GET("/:siteSlug/data", h.GetSiteData)

		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
			auth.POST("/logout", h.Logout)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware(config))
		{
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
		}
	}

	return router
}
