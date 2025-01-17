package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/FelipeSoft/home-broker-order-service/internal/application/grpc/service"
	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	grpcServerListening, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error on TCP listening: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	marketOrderService := service.NewMarketOrderService()
	proto.RegisterMarketOrderServiceServer(grpcServer, marketOrderService)

	go func() {
		log.Printf("The gRPC Server is listening on %s", grpcServerListening.Addr())
		if err := grpcServer.Serve(grpcServerListening); err != nil || err != grpc.ErrServerStopped {
			log.Fatalf("The gRPC server terminated unexpectedly: %v", err)
		}
	}()

	go func() {
		<-ctx.Done()
		log.Println("Shutting down gracefully")
		grpcServer.GracefulStop()
		log.Println("The gRPC server stopped")
		os.Exit(0)
	}()

	select {}
}