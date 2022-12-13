package product

import (
	"context"

	"github.com/devminnu/learn-rest/product/internal/app/model"
	"github.com/devminnu/learn-rest/product/internal/app/repository"
	"github.com/devminnu/learn-rest/product/internal/app/service"
	"github.com/google/uuid"
)

type productService struct {
	service.ProductService
	productRepository repository.ProductRepository
}

func New(productRepository repository.ProductRepository) service.ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) GetProduct(ctx context.Context, productID *model.ProductID) (product *model.Product, err error) {
	product, err = ps.productRepository.GetProduct(ctx, productID)
	if err != nil {
		return
	}

	return
}

func (ps *productService) AddProduct(ctx context.Context, product *model.Product) (productID *model.ProductID, err error) {
	productID = &model.ProductID{
		ID: uuid.New().String(),
	}
	product.ProductID = productID

	productID, err = ps.productRepository.AddProduct(ctx, product)
	if err != nil {
		return
	}

	return
}
