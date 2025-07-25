// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc1
// source: binance.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Binance-specific trade payload
type BinanceTrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType    string  `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`  // always "trade"
	EventTime    int64   `protobuf:"varint,2,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"` // event time in epoch ms
	Symbol       string  `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`                         // trading pair, e.g., "BTCUSDT"
	TradeTime    int64   `protobuf:"varint,4,opt,name=trade_time,json=tradeTime,proto3" json:"trade_time,omitempty"` // trade time in epoch ms
	IsBuyerMaker bool    `protobuf:"varint,5,opt,name=is_buyer_maker,json=isBuyerMaker,proto3" json:"is_buyer_maker,omitempty"`
	TradeId      string  `protobuf:"bytes,6,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	Price        float64 `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	Quantity     float64 `protobuf:"fixed64,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *BinanceTrade) Reset() {
	*x = BinanceTrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinanceTrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinanceTrade) ProtoMessage() {}

func (x *BinanceTrade) ProtoReflect() protoreflect.Message {
	mi := &file_binance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinanceTrade.ProtoReflect.Descriptor instead.
func (*BinanceTrade) Descriptor() ([]byte, []int) {
	return file_binance_proto_rawDescGZIP(), []int{0}
}

func (x *BinanceTrade) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *BinanceTrade) GetEventTime() int64 {
	if x != nil {
		return x.EventTime
	}
	return 0
}

func (x *BinanceTrade) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *BinanceTrade) GetTradeTime() int64 {
	if x != nil {
		return x.TradeTime
	}
	return 0
}

func (x *BinanceTrade) GetIsBuyerMaker() bool {
	if x != nil {
		return x.IsBuyerMaker
	}
	return false
}

func (x *BinanceTrade) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

func (x *BinanceTrade) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BinanceTrade) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// Wrapper tying raw Binance payload to canonical model
type BinanceMarketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw       *BinanceTrade `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Canonical *MarketData   `protobuf:"bytes,2,opt,name=canonical,proto3" json:"canonical,omitempty"` // populate fields from raw, symbol requires no conversion
}

func (x *BinanceMarketData) Reset() {
	*x = BinanceMarketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinanceMarketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinanceMarketData) ProtoMessage() {}

func (x *BinanceMarketData) ProtoReflect() protoreflect.Message {
	mi := &file_binance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinanceMarketData.ProtoReflect.Descriptor instead.
func (*BinanceMarketData) Descriptor() ([]byte, []int) {
	return file_binance_proto_rawDescGZIP(), []int{1}
}

func (x *BinanceMarketData) GetRaw() *BinanceTrade {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *BinanceMarketData) GetCanonical() *MarketData {
	if x != nil {
		return x.Canonical
	}
	return nil
}

var File_binance_proto protoreflect.FileDescriptor

var file_binance_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x42, 0x75, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x6b, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x6e,
	0x0a, 0x11, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x30, 0x0a, 0x09,
	0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x0d,
	0x5a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_binance_proto_rawDescOnce sync.Once
	file_binance_proto_rawDescData = file_binance_proto_rawDesc
)

func file_binance_proto_rawDescGZIP() []byte {
	file_binance_proto_rawDescOnce.Do(func() {
		file_binance_proto_rawDescData = protoimpl.X.CompressGZIP(file_binance_proto_rawDescData)
	})
	return file_binance_proto_rawDescData
}

var file_binance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_binance_proto_goTypes = []interface{}{
	(*BinanceTrade)(nil),      // 0: binance.BinanceTrade
	(*BinanceMarketData)(nil), // 1: binance.BinanceMarketData
	(*MarketData)(nil),        // 2: market.MarketData
}
var file_binance_proto_depIdxs = []int32{
	0, // 0: binance.BinanceMarketData.raw:type_name -> binance.BinanceTrade
	2, // 1: binance.BinanceMarketData.canonical:type_name -> market.MarketData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_binance_proto_init() }
func file_binance_proto_init() {
	if File_binance_proto != nil {
		return
	}
	file_trade_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_binance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinanceTrade); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_binance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinanceMarketData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_binance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_binance_proto_goTypes,
		DependencyIndexes: file_binance_proto_depIdxs,
		MessageInfos:      file_binance_proto_msgTypes,
	}.Build()
	File_binance_proto = out.File
	file_binance_proto_rawDesc = nil
	file_binance_proto_goTypes = nil
	file_binance_proto_depIdxs = nil
}
