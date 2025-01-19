package service

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/FelipeSoft/home-broker-order-service/internal/domain"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/akafka"
	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/limit_order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LimitOrderService struct {
	proto.UnimplementedLimitOrderServiceServer
	kafkaServers string
	wg           *sync.WaitGroup
}

func NewLimitOrderService(kafkaServers string, wg *sync.WaitGroup) *LimitOrderService {
	return &LimitOrderService{
		kafkaServers: kafkaServers,
		wg:           wg,
	}
}

func (s *LimitOrderService) BuyOrderByLimitValue(ctx context.Context, req *proto.BuyOrderByLimitValueRequest) (*emptypb.Empty, error) {
	ticker := req.GetTicker()
	quantity := float64(req.GetQuantity())
	minPrice := float64(req.GetMinPrice())
	maxPrice := float64(req.GetMaxPrice())

	order, err := domain.NewOrder(ticker, quantity, minPrice, &minPrice, &maxPrice, domain.OrderTypeBuy)
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
		akafka.Produce("matching_engine_topic", s.kafkaServers, orderSerialized)
	}()

	return &emptypb.Empty{}, nil
}

func (s *LimitOrderService) SellOrderByLimitValue(ctx context.Context, req *proto.SellOrderByLimitValueRequest) (*emptypb.Empty, error) {
	ticker := req.GetTicker()
	quantity := float64(req.GetQuantity())
	minPrice := float64(req.GetMinPrice())
	maxPrice := float64(req.GetMaxPrice())

	order, err := domain.NewOrder(ticker, quantity, maxPrice, &minPrice, &maxPrice, domain.OrderTypeSell)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Sell order error: %s", err.Error())
	}

	orderSerialized, err := json.Marshal(order)
	if err != nil {
		log.Printf("Error on sell order serializing: %s", err.Error())
	}

	go akafka.Produce("matching_engine_topic", s.kafkaServers, orderSerialized)
	return &emptypb.Empty{}, nil
}
