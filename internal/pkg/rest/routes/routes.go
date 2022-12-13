package routes

import (
	"context"
	"net/http"

	"github.com/devminnu/learn-rest/product/internal/app/controller/rest/product"
	postgresrepository "github.com/devminnu/learn-rest/product/internal/app/repository/postgres"
	productservice "github.com/devminnu/learn-rest/product/internal/app/service/product"
	"github.com/devminnu/learn-rest/product/pkg/db/postgres"
	"github.com/gin-gonic/gin"
)

func LoadRestHandlers(ctx context.Context, r *gin.Engine) {
	registerPingHandler(r)

	service := productservice.New(postgresrepository.New(postgres.Connect(ctx)))

	r.POST("/getProduct", product.New(service).GetProduct)
	r.POST("/addProduct", product.New(service).AddProduct)
}

func registerPingHandler(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, `{"ping":"ok"}`)
	})
}
