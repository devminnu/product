package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Host string `env:"GRPC_HOST,default=localhost"`
	Port string `env:"GRPC_PORT,default=50051"`
}

func Run(ctx context.Context, registerServices ...func(ctx context.Context, grpcServer *grpc.Server)) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	// register grpc services
	for _, registerService := range registerServices {
		registerService(ctx, grpcServer)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
