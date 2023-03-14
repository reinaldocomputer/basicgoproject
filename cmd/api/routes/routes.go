package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/reinaldocomputer/basicgoproject/cmd/api/handlers"
)

func API() {
	r := gin.Default()
	r.GET("/ping", handlers.PingHandler)

	// Basic Handler
	basicCRUDGroup := r.Group("/basic")
	basicCRUDGroup.DELETE("/:id", handlers.DeleteBasicByID)
	basicCRUDGroup.GET("", handlers.GetBasicAll)
	basicCRUDGroup.GET("/:id", handlers.GetBasicByID)
	basicCRUDGroup.POST("", handlers.InsertBasic)
	basicCRUDGroup.PUT("/:id", handlers.UpdateBasicByID)

	// Initialize
	r.Run("localhost:8080")
}
