// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.2
// source: market_order_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BuyOrderByMarketValueRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ticker        string                 `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Price         float32                `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      float32                `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuyOrderByMarketValueRequest) Reset() {
	*x = BuyOrderByMarketValueRequest{}
	mi := &file_market_order_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuyOrderByMarketValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrderByMarketValueRequest) ProtoMessage() {}

func (x *BuyOrderByMarketValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_order_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyOrderByMarketValueRequest.ProtoReflect.Descriptor instead.
func (*BuyOrderByMarketValueRequest) Descriptor() ([]byte, []int) {
	return file_market_order_service_proto_rawDescGZIP(), []int{0}
}

func (x *BuyOrderByMarketValueRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *BuyOrderByMarketValueRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BuyOrderByMarketValueRequest) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type SellOrderByMarketValueRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ticker        string                 `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Price         float32                `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      float32                `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SellOrderByMarketValueRequest) Reset() {
	*x = SellOrderByMarketValueRequest{}
	mi := &file_market_order_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SellOrderByMarketValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellOrderByMarketValueRequest) ProtoMessage() {}

func (x *SellOrderByMarketValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_order_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellOrderByMarketValueRequest.ProtoReflect.Descriptor instead.
func (*SellOrderByMarketValueRequest) Descriptor() ([]byte, []int) {
	return file_market_order_service_proto_rawDescGZIP(), []int{1}
}

func (x *SellOrderByMarketValueRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *SellOrderByMarketValueRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SellOrderByMarketValueRequest) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_market_order_service_proto protoreflect.FileDescriptor

var file_market_order_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x1c, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x69, 0x0a, 0x1d, 0x53, 0x65, 0x6c,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x32, 0xe4, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x15, 0x42,
	0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x54, 0x5a, 0x52, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x6c, 0x69, 0x70, 0x65,
	0x53, 0x6f, 0x66, 0x74, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_market_order_service_proto_rawDescOnce sync.Once
	file_market_order_service_proto_rawDescData = file_market_order_service_proto_rawDesc
)

func file_market_order_service_proto_rawDescGZIP() []byte {
	file_market_order_service_proto_rawDescOnce.Do(func() {
		file_market_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_market_order_service_proto_rawDescData)
	})
	return file_market_order_service_proto_rawDescData
}

var file_market_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_market_order_service_proto_goTypes = []any{
	(*BuyOrderByMarketValueRequest)(nil),  // 0: market_order_service.BuyOrderByMarketValueRequest
	(*SellOrderByMarketValueRequest)(nil), // 1: market_order_service.SellOrderByMarketValueRequest
	(*emptypb.Empty)(nil),                 // 2: google.protobuf.Empty
}
var file_market_order_service_proto_depIdxs = []int32{
	0, // 0: market_order_service.MarketOrderService.BuyOrderByMarketValue:input_type -> market_order_service.BuyOrderByMarketValueRequest
	1, // 1: market_order_service.MarketOrderService.SellOrderByMarketValue:input_type -> market_order_service.SellOrderByMarketValueRequest
	2, // 2: market_order_service.MarketOrderService.BuyOrderByMarketValue:output_type -> google.protobuf.Empty
	2, // 3: market_order_service.MarketOrderService.SellOrderByMarketValue:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_market_order_service_proto_init() }
func file_market_order_service_proto_init() {
	if File_market_order_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_order_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_market_order_service_proto_goTypes,
		DependencyIndexes: file_market_order_service_proto_depIdxs,
		MessageInfos:      file_market_order_service_proto_msgTypes,
	}.Build()
	File_market_order_service_proto = out.File
	file_market_order_service_proto_rawDesc = nil
	file_market_order_service_proto_goTypes = nil
	file_market_order_service_proto_depIdxs = nil
}
