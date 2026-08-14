package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Reponse struct {
	Message string
	Status  bool
}

func TestHandlers(c *gin.Context) {
	c.JSON(http.StatusOK, Reponse{Message: "test check for end point", Status: true})
}
