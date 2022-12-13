package postgres

import (
	"context"
	"database/sql"

	"github.com/devminnu/learn-rest/product/internal/app/model"
	"github.com/devminnu/learn-rest/product/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	GetProduct = "select id, name from product where id=$1;"
	AddProduct = "insert into product(id,name) values($1,$2);"
)

type productRepository struct {
	repository.ProductRepository
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

func (pr *productRepository) GetProduct(ctx context.Context, p *model.ProductID) (product *model.Product, err error) {
	product = new(model.Product)
	v := &struct {
		ID   string
		Name string
	}{}
	err = pr.db.QueryRowx(GetProduct, p.ID).StructScan(v)
	if err != nil {
		if err == sql.ErrNoRows {
			return &model.Product{}, nil
		}
		log.Error().Err(err).Msg("error retreiving product from db")

		return
	}
	product = &model.Product{
		ProductID: &model.ProductID{ID: v.ID},
		Name:      v.Name,
	}

	return
}

func (pr *productRepository) AddProduct(ctx context.Context, p *model.Product) (productID *model.ProductID, err error) {
	_, err = pr.db.Exec(AddProduct, p.ProductID.ID, p.Name)
	if err != nil {
		log.Error().Err(err).Msg("error adding product in db")
	}

	return
}
