package main

import (
	"context"

	"github.com/devminnu/learn-rest/product/internal/pkg/grpc/controllers"
	grpc_server "github.com/devminnu/learn-rest/product/pkg/grpc/server"
	"github.com/devminnu/learn-rest/product/pkg/logger"
)

func init() {
	ctx := context.Background()
	logger.Init(ctx)
}

func main() {
	ctx := context.Background()
	// rest_server.Run(ctx, routes.LoadRestHandlers)
	grpc_server.Run(ctx, controllers.RegisterGRPCServices()...)
}
