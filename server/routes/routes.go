package routes

import (
	"gin-atomy/server/service/products"
	"github.com/gin-gonic/gin"
)

// Engine - engine
func Engine() *gin.Engine {
	r := gin.Default()

	// products
	r.GET("/api/atomy/products", products.GetProducts())
	r.POST("/callback", products.GetAtomyProducts())

	return r
}
