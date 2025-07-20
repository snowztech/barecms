package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCollection(c *gin.Context) {
	var req models.CreateCollectionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateCollection(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Collection created!"})
}

func (h *Handler) GetCollection(c *gin.Context) {
	id := c.Param("id")

	collection, err := h.Service.GetCollectionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, collection)
}

func (h *Handler) GetCollectionsBySiteID(c *gin.Context) {
	siteID := c.Param("id")

	collections, err := h.Service.GetCollectionsBySiteID(siteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"collections": collections})
}

func (h *Handler) DeleteCollection(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteCollection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collection deleted!"})
}
