package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meekailkhan/quick-book/internals/user"
)

// InventoryHandler holds dependencies for inventory HTTP handlers.
type InventoryHandler struct {
	service *user.InventoryService
}

// NewInventoryHandler creates a new InventoryHandler.
func NewInventoryHandler(service *user.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

// GetAllInventories godoc
// GET /api/v1/user/inventories/get-all
// Query params:
//
//	page     int    (default: 1)
//	category string (optional)
func (h *InventoryHandler) GetAllInventories(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	category := c.Query("category") // empty string = no filter

	result, err := h.service.GetAllInventories(c.Request.Context(), user.GetAllInventoriesParams{
		Page:     page,
		Category: category,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "failed to fetch inventories",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "inventories fetched successfully",
		"data":    result,
	})
}
