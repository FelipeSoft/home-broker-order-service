package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FelipeSoft/home-broker-order-service/internal/application/grpc/service"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/akafka"
	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	marketOrderService := service.NewMarketOrderService("localhost:9092")
	proto.RegisterMarketOrderServiceServer(grpcServer, marketOrderService)

	kafkaMessagesChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"matching_engine_topic"}, "localhost:9092", "matching_engine_consumer_group", kafkaMessagesChan)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-kafkaMessagesChan:
				if msg == nil {
					log.Println("Kafka channel closed, stopping consumer")
					return
				}
				log.Printf("Received Kafka message: %s", string(msg.Value))
			case <-ctx.Done():
				log.Println("Stopping Kafka consumer due to context cancellation")
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("The gRPC Server is listening on %s", grpcServerListening.Addr())
		if err := grpcServer.Serve(grpcServerListening); err != nil || err != grpc.ErrServerStopped {
			log.Fatalf("The gRPC server terminated unexpectedly: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully")

	grpcServer.GracefulStop()
	log.Println("The gRPC server stopped")

	wg.Wait()
	log.Println("Application shutdown complete")
}
