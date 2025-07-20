package handlers

import (
	"barecms/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEntry(c *gin.Context) {
	var request models.CreateEntryRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateEntry(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Entry created!"})
}

func (h *Handler) GetEntry(c *gin.Context) {
	id := c.Param("id")

	entry, err := h.Service.GetEntryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func (h *Handler) GetCollectionEntries(c *gin.Context) {
	id := c.Param("id")

	entries, err := h.Service.GetEntriesByCollectionID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func (h *Handler) DeleteEntry(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteEntry(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted!"})
}
