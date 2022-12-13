package mongo

import (
	"context"

	"github.com/devminnu/learn-rest/product/internal/app/model"
	"github.com/devminnu/learn-rest/product/internal/app/repository"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	repository.ProductRepository
	db *mongo.Database
}

func New(db *mongo.Database) repository.ProductRepository {
	return &productRepository{db: db}
}

func (pr *productRepository) GetProduct(ctx context.Context, p *model.ProductID) (product *model.Product, err error) {

	return
}

func (pr *productRepository) AddProduct(ctx context.Context, p *model.Product) (productID *model.ProductID, err error) {
	_, err = pr.db.Collection("product").InsertOne(ctx, p)
	if err != nil {
		log.Error().Err(err).Msg("error inserting product")

		return
	}

	return
}
