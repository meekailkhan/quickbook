package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/meekailkhan/quick-book/internals/handlers"
	"github.com/meekailkhan/quick-book/internals/user"
)

// Add new handlers here as the app grows.
type AppHandlers struct {
	Inventory *handlers.InventoryHandler
}

func newAppHandlers(db *sql.DB) *AppHandlers {
	inventoryService := user.NewInventoryService(db)

	return &AppHandlers{
		Inventory: handlers.NewInventoryHandler(inventoryService),
	}
}

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	h := newAppHandlers(db)

	// misc
	r.GET("/test1", handlers.TestHandlers)

	// user group
	userGroup := r.Group("/api/v1/user")
	{
		userGroup.GET("/inventories/get-all", h.Inventory.GetAllInventories)
	}

	return r
}
