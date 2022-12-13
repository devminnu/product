package rest

import (
	"github.com/gin-gonic/gin"
)

type API interface {
	Product
}

type Product interface {
	GetProduct(c *gin.Context)
	AddProduct(c *gin.Context)
}
