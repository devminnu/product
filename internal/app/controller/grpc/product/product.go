package product

import (
	"context"

	pb_product "github.com/devminnu/learn-rest/product/api/grpc/protogen/product"
)

type productServer struct {
	pb_product.UnimplementedProductServer
}

func New() pb_product.ProductServer {
	return &productServer{}
}

func (p *productServer) AddProduct(ctx context.Context, in *pb_product.AddProductRequest) (AddProductResponse *pb_product.AddProductResponse, err error) {
	AddProductResponse = &pb_product.AddProductResponse{
		Id: "123",
	}
	return
}
