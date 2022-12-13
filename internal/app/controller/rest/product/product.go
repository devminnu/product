package product

import (
	"net/http"

	"github.com/devminnu/learn-rest/product/internal/app/controller/rest"
	"github.com/devminnu/learn-rest/product/internal/app/model"
	"github.com/devminnu/learn-rest/product/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type product struct {
	rest.Product
	service service.ProductService
}

func New(service service.ProductService) rest.Product {
	return &product{service: service}
}

func (p *product) GetProduct(c *gin.Context) {
	productID := new(model.ProductID)
	err := c.ShouldBindJSON(productID)
	if err != nil {
		log.Error().Err(err).Msg("bind error ProductID")
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}
	product, err := p.service.GetProduct(c.Request.Context(), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *product) AddProduct(c *gin.Context) {
	product := new(model.Product)
	err := c.ShouldBindJSON(product)
	if err != nil {
		log.Error().Err(err).Msg("bind error ProductID")
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}
	productID, err := p.service.AddProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, productID)
}
