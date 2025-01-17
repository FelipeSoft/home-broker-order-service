// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: market_order_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MarketOrderService_BuyOrderByMarketValue_FullMethodName  = "/market_order_service.MarketOrderService/BuyOrderByMarketValue"
	MarketOrderService_SellOrderByMarketValue_FullMethodName = "/market_order_service.MarketOrderService/SellOrderByMarketValue"
)

// MarketOrderServiceClient is the client API for MarketOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketOrderServiceClient interface {
	BuyOrderByMarketValue(ctx context.Context, in *BuyOrderByMarketValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SellOrderByMarketValue(ctx context.Context, in *SellOrderByMarketValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type marketOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketOrderServiceClient(cc grpc.ClientConnInterface) MarketOrderServiceClient {
	return &marketOrderServiceClient{cc}
}

func (c *marketOrderServiceClient) BuyOrderByMarketValue(ctx context.Context, in *BuyOrderByMarketValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MarketOrderService_BuyOrderByMarketValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketOrderServiceClient) SellOrderByMarketValue(ctx context.Context, in *SellOrderByMarketValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MarketOrderService_SellOrderByMarketValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketOrderServiceServer is the server API for MarketOrderService service.
// All implementations must embed UnimplementedMarketOrderServiceServer
// for forward compatibility.
type MarketOrderServiceServer interface {
	BuyOrderByMarketValue(context.Context, *BuyOrderByMarketValueRequest) (*emptypb.Empty, error)
	SellOrderByMarketValue(context.Context, *SellOrderByMarketValueRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMarketOrderServiceServer()
}

// UnimplementedMarketOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMarketOrderServiceServer struct{}

func (UnimplementedMarketOrderServiceServer) BuyOrderByMarketValue(context.Context, *BuyOrderByMarketValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyOrderByMarketValue not implemented")
}
func (UnimplementedMarketOrderServiceServer) SellOrderByMarketValue(context.Context, *SellOrderByMarketValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrderByMarketValue not implemented")
}
func (UnimplementedMarketOrderServiceServer) mustEmbedUnimplementedMarketOrderServiceServer() {}
func (UnimplementedMarketOrderServiceServer) testEmbeddedByValue()                            {}

// UnsafeMarketOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketOrderServiceServer will
// result in compilation errors.
type UnsafeMarketOrderServiceServer interface {
	mustEmbedUnimplementedMarketOrderServiceServer()
}

func RegisterMarketOrderServiceServer(s grpc.ServiceRegistrar, srv MarketOrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedMarketOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MarketOrderService_ServiceDesc, srv)
}

func _MarketOrderService_BuyOrderByMarketValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyOrderByMarketValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketOrderServiceServer).BuyOrderByMarketValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketOrderService_BuyOrderByMarketValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketOrderServiceServer).BuyOrderByMarketValue(ctx, req.(*BuyOrderByMarketValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketOrderService_SellOrderByMarketValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellOrderByMarketValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketOrderServiceServer).SellOrderByMarketValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketOrderService_SellOrderByMarketValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketOrderServiceServer).SellOrderByMarketValue(ctx, req.(*SellOrderByMarketValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketOrderService_ServiceDesc is the grpc.ServiceDesc for MarketOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market_order_service.MarketOrderService",
	HandlerType: (*MarketOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyOrderByMarketValue",
			Handler:    _MarketOrderService_BuyOrderByMarketValue_Handler,
		},
		{
			MethodName: "SellOrderByMarketValue",
			Handler:    _MarketOrderService_SellOrderByMarketValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market_order_service.proto",
}
