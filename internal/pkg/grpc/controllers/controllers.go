package controllers

import (
	"context"

	pb_product "github.com/devminnu/learn-rest/product/api/grpc/protogen/product"
	"github.com/devminnu/learn-rest/product/internal/app/controller/grpc/product"
	"google.golang.org/grpc"
)

func RegisterGRPCServices() (grpcServices []func(context.Context, *grpc.Server)) {
	grpcServices = append(grpcServices, registerProductService)

	return
}

func registerProductService(ctx context.Context, grpcServer *grpc.Server) {
	productServer := product.New()
	pb_product.RegisterProductServer(grpcServer, productServer)
}
