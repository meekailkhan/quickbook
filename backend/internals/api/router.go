package api

import (
	"github.com/gin-gonic/gin"
	"github.com/meekailkhan/quick-book/internals/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// v1 := r.Group("/api/v1")
	r.GET("/test1", handlers.TestHandlers)

	return r
}
