package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/FelipeSoft/home-broker-order-service/internal/domain"
	"github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/akafka"
	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MarketOrderService struct {
	proto.UnimplementedMarketOrderServiceServer
	kafkaServers    string
	wg              *sync.WaitGroup
	orderRepository domain.OrderRepository
}

func NewMarketOrderService(kafkaServers string, wg *sync.WaitGroup, orderRepository domain.OrderRepository) *MarketOrderService {
	return &MarketOrderService{
		kafkaServers:    kafkaServers,
		wg:              wg,
		orderRepository: orderRepository,
	}
}

func (s *MarketOrderService) BuyOrderByMarketValue(ctx context.Context, req *proto.BuyOrderByMarketValueRequest) (*emptypb.Empty, error) {
	ticker := req.GetTicker()
	quantity := req.Quantity
	price := req.Price
	userId := fmt.Sprint(req.GetUserId())
	log.Printf("Ticker: %s; Price: %f; Quantity: %f", ticker, price, quantity)

	order, err := domain.NewOrder(ticker, quantity, price, nil, nil, domain.OrderTypeBuy, userId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Buy order error: %s", err.Error())
	}

	orderId, err := s.orderRepository.SaveOrder(*order)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Order saving error: %s", err.Error())
	}

	order.ID = orderId

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
	quantity := req.GetQuantity()
	price := req.GetPrice()
	userId := fmt.Sprint(req.GetUserId())

	order, err := domain.NewOrder(ticker, quantity, price, nil, nil, domain.OrderTypeSell, userId)
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
