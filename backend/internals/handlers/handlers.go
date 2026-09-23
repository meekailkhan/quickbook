package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meekailkhan/quick-book/internals/user"
)

type Reponse struct {
	Message string
	Status  bool
}

func TestHandlers(c *gin.Context) {
	c.JSON(http.StatusOK, Reponse{Message: "test check for end point", Status: true})
}

// Add new handlers here as the app grows.
type AppHandlers struct {
	Inventory *InventoryHandler
}

func NewAppHandlers(db *sql.DB) *AppHandlers {
	inventoryService := user.NewInventoryService(db)

	return &AppHandlers{
		Inventory: NewInventoryHandler(inventoryService),
	}
}
