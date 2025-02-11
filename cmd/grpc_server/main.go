package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	_ "github.com/go-sql-driver/mysql"
	"github.com/FelipeSoft/home-broker-order-service/internal/application/grpc/service"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/akafka"
	proto_limit_order "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/limit_order"
	proto_market_order "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/mysql"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatalf("Error during the .env file loading: %s", err.Error())
	}
	wg := &sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	grpcServerListening, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error on TCP listening: %s", err.Error())
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_SERVER"))
	if err != nil {
		log.Fatalf("Error during MYSQL Connection: %s", err.Error())
	}
	
	orderRepositoryMySql := mysql.NewOrderRepositoryMySql(db)

	grpcServer := grpc.NewServer()
	marketOrderService := service.NewMarketOrderService("localhost:9092", wg, orderRepositoryMySql)
	proto_market_order.RegisterMarketOrderServiceServer(grpcServer, marketOrderService)

	limitOrderService := service.NewLimitOrderService("localhost:9092", wg)
	proto_limit_order.RegisterLimitOrderServiceServer(grpcServer, limitOrderService)

	kafkaMessagesChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"matching_engine_topic"}, "localhost:9092", "matching_engine_consumer_group", kafkaMessagesChan)

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
