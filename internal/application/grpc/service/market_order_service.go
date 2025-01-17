package service

import (
	"context"

	proto "github.com/FelipeSoft/home-broker-order-service/internal/infrastructure/grpc/proto/market_order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MarketOrderService struct {
	proto.UnimplementedMarketOrderServiceServer
}

func NewMarketOrderService() *MarketOrderService {
	return &MarketOrderService{}
}

func (s *MarketOrderService) BuyOrderByMarketValue(ctx context.Context, req *proto.BuyOrderByMarketValueRequest) (*emptypb.Empty, error) {
    if req.Ticker == "" {
        return nil, status.Errorf(codes.InvalidArgument, "ticker is required")
    }
    if req.Price <= 0 {
        return nil, status.Errorf(codes.InvalidArgument, "price must be greater than 0")
    }
    if req.Quantity <= 0 {
        return nil, status.Errorf(codes.InvalidArgument, "quantity must be greater than 0")
    }

    // Process the request
    return &emptypb.Empty{}, nil
}

func (s *MarketOrderService) SellOrderByMarketValue(ctx context.Context, message *proto.SellOrderByMarketValueRequest) (*emptypb.Empty, error) {
	// process request
	return &emptypb.Empty{}, nil
}