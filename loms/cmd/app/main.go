package main

import (
	"log"
	"net"
	"route256/libs/middleware"
	"route256/loms/internal/app/loms"
	"route256/loms/internal/config"
	"route256/loms/internal/usecase"
	desc "route256/loms/pkg/loms"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()

	useCaseInstance := usecase.New()
	serviceInstance := loms.New(useCaseInstance)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				middleware.LoggingInterceptor,
			),
		),
	)

	reflection.Register(s)
	desc.RegisterLomsServer(s, serviceInstance)

	log.Printf("start \"loms\" checkout at %s\n", config.Instance.Services.Loms)
	lis, err := net.Listen("tcp", config.Instance.Services.Loms)
	if err != nil {
		log.Fatalf("create tcp listener: %v", err)
	}
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
