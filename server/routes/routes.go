package routes

import (
	"github.com/gin-gonic/gin"
	"gin-atomy/server/service/products"
)

func Engine() *gin.Engine {
	r := gin.Default()

	// products
	r.GET("/api/atomy/products", products.GetProducts())
	r.POST("/callback", products.GetAtomyProducts())

	return r
}
