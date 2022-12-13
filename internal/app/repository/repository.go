package repository

import (
	"context"

	"github.com/devminnu/learn-rest/product/internal/app/model"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, productID *model.ProductID) (product *model.Product, err error)
	AddProduct(ctx context.Context, product *model.Product) (productID *model.ProductID, err error)
}
