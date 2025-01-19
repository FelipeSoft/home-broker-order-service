package service

import (
	"context"
	"encoding/json"
	"github.com/FelipeSoft/home-broker-order-service/internal/domain"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/akafka"
	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"sync"
)

type MarketOrderService struct {
	proto.UnimplementedMarketOrderServiceServer
	kafkaServers string
	wg           *sync.WaitGroup
}

func NewMarketOrderService(kafkaServers string, wg *sync.WaitGroup) *MarketOrderService {
	return &MarketOrderService{
		kafkaServers: kafkaServers,
		wg:           wg,
	}
}

func (s *MarketOrderService) BuyOrderByMarketValue(ctx context.Context, req *proto.BuyOrderByMarketValueRequest) (*emptypb.Empty, error) {
	ticker := req.GetTicker()
	quantity := float64(req.GetQuantity())
	price := float64(req.GetPrice())

	order, err := domain.NewOrder(ticker, quantity, price, nil, nil, domain.OrderTypeBuy)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Buy order error: %s", err.Error())
	}

	orderSerialized, err := json.Marshal(order)
	if err != nil {
		log.Printf("Error on buy order serializing: %s", err.Error())
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		akafka.Produce("matching_engine_topic", s.kafkaServers, []byte(orderSerialized))
	}()

	return &emptypb.Empty{}, nil
}

func (s *MarketOrderService) SellOrderByMarketValue(ctx context.Context, req *proto.SellOrderByMarketValueRequest) (*emptypb.Empty, error) {
	ticker := req.GetTicker()
	quantity := float64(req.GetQuantity())
	price := float64(req.GetPrice())

	order, err := domain.NewOrder(ticker, quantity, price, nil, nil, domain.OrderTypeSell)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Sell order error: %s", err.Error())
	}

	orderSerialized, err := json.Marshal(order)
	if err != nil {
		log.Printf("Error on sell order serializing: %s", err.Error())
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		akafka.Produce("matching_engine_topic", s.kafkaServers, []byte(orderSerialized))
	}()

	return &emptypb.Empty{}, nil
}
