package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/FelipeSoft/home-broker-order-service/internal/domain"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/akafka"
	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MarketOrderService struct {
	proto.UnimplementedMarketOrderServiceServer
	kafkaServers string
}

func NewMarketOrderService(kafkaServers string) *MarketOrderService {
	return &MarketOrderService{
		kafkaServers: kafkaServers,
	}
}

func (s *MarketOrderService) BuyOrderByMarketValue(ctx context.Context, req *proto.BuyOrderByMarketValueRequest) (*emptypb.Empty, error) {
	ticker := req.GetTicker()
	quantity := req.GetQuantity()
	price := req.GetPrice()

	order, err := domain.NewOrder(ticker, quantity, price)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Buy Order by Market Value Error: %s", err.Error())
	}

	marketOrder := domain.NewMarketOrder(order)
	fmt.Println(marketOrder)

	marketOrderSerialized, err := json.Marshal(marketOrder)
	if err != nil {
		log.Printf("Error on market order serializing: %s", err.Error())
	}

	// Send market order to Apache Kafka topic responsible for Matching Engine method.
	go akafka.Produce("matching_engine_topic", s.kafkaServers, []byte(marketOrderSerialized))

	return &emptypb.Empty{}, nil
}

func (s *MarketOrderService) SellOrderByMarketValue(ctx context.Context, message *proto.SellOrderByMarketValueRequest) (*emptypb.Empty, error) {
	// process request
	return &emptypb.Empty{}, nil
}
