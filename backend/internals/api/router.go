package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/meekailkhan/quick-book/internals/handlers"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	h := handlers.NewAppHandlers(db)

	// misc
	r.GET("/test1", handlers.TestHandlers)

	// user group
	userGroup := r.Group("/api/v1/user")
	{
		userGroup.GET("/inventories/get-all", h.Inventory.GetAllInventories)
	}

	return r
}
