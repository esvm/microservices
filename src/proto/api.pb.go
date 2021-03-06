// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Metric struct {
	Average              float32  `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
	Median               float32  `protobuf:"fixed32,2,opt,name=median,proto3" json:"median,omitempty"`
	Variance             float32  `protobuf:"fixed32,3,opt,name=variance,proto3" json:"variance,omitempty"`
	StandardDeviation    float32  `protobuf:"fixed32,4,opt,name=standard_deviation,json=standardDeviation,proto3" json:"standard_deviation,omitempty"`
	Mode                 float32  `protobuf:"fixed32,5,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetAverage() float32 {
	if m != nil {
		return m.Average
	}
	return 0
}

func (m *Metric) GetMedian() float32 {
	if m != nil {
		return m.Median
	}
	return 0
}

func (m *Metric) GetVariance() float32 {
	if m != nil {
		return m.Variance
	}
	return 0
}

func (m *Metric) GetStandardDeviation() float32 {
	if m != nil {
		return m.StandardDeviation
	}
	return 0
}

func (m *Metric) GetMode() float32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type InsertItemRequest struct {
	TopicName            string    `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	Metrics              []*Metric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InsertItemRequest) Reset()         { *m = InsertItemRequest{} }
func (m *InsertItemRequest) String() string { return proto.CompactTextString(m) }
func (*InsertItemRequest) ProtoMessage()    {}
func (*InsertItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{1}
}

func (m *InsertItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertItemRequest.Unmarshal(m, b)
}
func (m *InsertItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertItemRequest.Marshal(b, m, deterministic)
}
func (m *InsertItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertItemRequest.Merge(m, src)
}
func (m *InsertItemRequest) XXX_Size() int {
	return xxx_messageInfo_InsertItemRequest.Size(m)
}
func (m *InsertItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertItemRequest proto.InternalMessageInfo

func (m *InsertItemRequest) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *InsertItemRequest) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type GetItemsRequest struct {
	TopicName            string   `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	MinutesAgo           int32    `protobuf:"varint,2,opt,name=minutes_ago,json=minutesAgo,proto3" json:"minutes_ago,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemsRequest) Reset()         { *m = GetItemsRequest{} }
func (m *GetItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemsRequest) ProtoMessage()    {}
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{2}
}

func (m *GetItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsRequest.Unmarshal(m, b)
}
func (m *GetItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsRequest.Marshal(b, m, deterministic)
}
func (m *GetItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsRequest.Merge(m, src)
}
func (m *GetItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemsRequest.Size(m)
}
func (m *GetItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsRequest proto.InternalMessageInfo

func (m *GetItemsRequest) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *GetItemsRequest) GetMinutesAgo() int32 {
	if m != nil {
		return m.MinutesAgo
	}
	return 0
}

type GetItemsResponse struct {
	Metrics              []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetItemsResponse) Reset()         { *m = GetItemsResponse{} }
func (m *GetItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetItemsResponse) ProtoMessage()    {}
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf0878b123623e2, []int{3}
}

func (m *GetItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsResponse.Unmarshal(m, b)
}
func (m *GetItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsResponse.Marshal(b, m, deterministic)
}
func (m *GetItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsResponse.Merge(m, src)
}
func (m *GetItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetItemsResponse.Size(m)
}
func (m *GetItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsResponse proto.InternalMessageInfo

func (m *GetItemsResponse) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Metric)(nil), "api.Metric")
	proto.RegisterType((*InsertItemRequest)(nil), "api.InsertItemRequest")
	proto.RegisterType((*GetItemsRequest)(nil), "api.GetItemsRequest")
	proto.RegisterType((*GetItemsResponse)(nil), "api.GetItemsResponse")
}

func init() { proto.RegisterFile("proto/api.proto", fileDescriptor_ecf0878b123623e2) }

var fileDescriptor_ecf0878b123623e2 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcf, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0xfd, 0xdd, 0xe9, 0xa1, 0x5f, 0x97, 0xaf, 0x25, 0xe4, 0xe3, 0xc3, 0x12, 0x10, 0x7a,
	0x31, 0x85, 0x7a, 0x10, 0xc1, 0x4b, 0xa5, 0x22, 0x3d, 0x28, 0x18, 0x4f, 0x7a, 0x29, 0xd3, 0x64,
	0x0c, 0x0b, 0xee, 0x6e, 0xdc, 0xdd, 0x06, 0xfc, 0x27, 0xbc, 0xfb, 0xdf, 0x4a, 0x37, 0x8d, 0xc5,
	0x8a, 0xe0, 0x69, 0x77, 0xde, 0x9b, 0x79, 0xbc, 0x79, 0x03, 0xfd, 0x5c, 0x2b, 0xab, 0xa6, 0x98,
	0xf3, 0xc8, 0xfd, 0x58, 0x1d, 0x73, 0x1e, 0xfc, 0xcb, 0x94, 0xca, 0x9e, 0x69, 0xea, 0xa0, 0xf5,
	0xe6, 0x69, 0x4a, 0x22, 0xb7, 0xaf, 0x65, 0x47, 0xf8, 0xee, 0x41, 0xeb, 0x86, 0xac, 0xe6, 0x09,
	0xf3, 0xa1, 0x8d, 0x05, 0x69, 0xcc, 0xc8, 0xf7, 0xc6, 0xde, 0xa4, 0x16, 0x57, 0x25, 0x1b, 0x41,
	0x4b, 0x50, 0xca, 0x51, 0xfa, 0x35, 0x47, 0xec, 0x2a, 0x16, 0x40, 0xa7, 0x40, 0xcd, 0x51, 0x26,
	0xe4, 0xd7, 0x1d, 0xf3, 0x59, 0xb3, 0x13, 0x60, 0xc6, 0xa2, 0x4c, 0x51, 0xa7, 0xab, 0x94, 0x0a,
	0x8e, 0x96, 0x2b, 0xe9, 0x37, 0x5c, 0xd7, 0xa0, 0x62, 0x16, 0x15, 0xc1, 0x18, 0x34, 0x84, 0x4a,
	0xc9, 0x6f, 0xba, 0x06, 0xf7, 0x0f, 0x1f, 0x60, 0xb0, 0x94, 0x86, 0xb4, 0x5d, 0x5a, 0x12, 0x31,
	0xbd, 0x6c, 0xc8, 0x58, 0xf6, 0x1f, 0xc0, 0xaa, 0x9c, 0x27, 0x2b, 0x89, 0xa2, 0x34, 0xda, 0x8d,
	0xbb, 0x0e, 0xb9, 0x45, 0x41, 0xec, 0x18, 0xda, 0xc2, 0xad, 0x63, 0xfc, 0xda, 0xb8, 0x3e, 0xe9,
	0xcd, 0x7a, 0xd1, 0x36, 0x8e, 0x72, 0xc5, 0xb8, 0xe2, 0xc2, 0x3b, 0xe8, 0x5f, 0x93, 0xd3, 0x35,
	0xbf, 0x14, 0x3e, 0x82, 0x9e, 0xe0, 0x72, 0x63, 0xc9, 0xac, 0x30, 0x53, 0x2e, 0x88, 0x66, 0x0c,
	0x3b, 0x68, 0x9e, 0xa9, 0xf0, 0x1c, 0xfe, 0xec, 0x25, 0x4d, 0xae, 0xa4, 0xf9, 0xe2, 0xc6, 0xfb,
	0xd9, 0xcd, 0xec, 0xcd, 0x83, 0xe1, 0x22, 0x11, 0xf3, 0x9c, 0x2f, 0xd0, 0xe2, 0x1a, 0x0d, 0xdd,
	0x93, 0x2e, 0x78, 0x42, 0xec, 0x02, 0x60, 0x1f, 0x01, 0x1b, 0xb9, 0xe9, 0x6f, 0x99, 0x04, 0xa3,
	0xa8, 0x3c, 0x71, 0x54, 0x9d, 0x38, 0xba, 0xda, 0x9e, 0x98, 0x9d, 0x41, 0xa7, 0xb2, 0xc4, 0xfe,
	0xba, 0xd9, 0x83, 0xa5, 0x83, 0xe1, 0x01, 0x5a, 0xfa, 0xbe, 0x6c, 0x3f, 0x36, 0x4b, 0xa9, 0x96,
	0x7b, 0x4e, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37, 0xc1, 0xf1, 0x97, 0x5a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DcmApiDatabaseServiceClient is the client API for DcmApiDatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DcmApiDatabaseServiceClient interface {
	InsertItem(ctx context.Context, in *InsertItemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error)
}

type dcmApiDatabaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewDcmApiDatabaseServiceClient(cc *grpc.ClientConn) DcmApiDatabaseServiceClient {
	return &dcmApiDatabaseServiceClient{cc}
}

func (c *dcmApiDatabaseServiceClient) InsertItem(ctx context.Context, in *InsertItemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DcmApiDatabaseService/InsertItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dcmApiDatabaseServiceClient) GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, "/api.DcmApiDatabaseService/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DcmApiDatabaseServiceServer is the server API for DcmApiDatabaseService service.
type DcmApiDatabaseServiceServer interface {
	InsertItem(context.Context, *InsertItemRequest) (*empty.Empty, error)
	GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error)
}

// UnimplementedDcmApiDatabaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDcmApiDatabaseServiceServer struct {
}

func (*UnimplementedDcmApiDatabaseServiceServer) InsertItem(ctx context.Context, req *InsertItemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertItem not implemented")
}
func (*UnimplementedDcmApiDatabaseServiceServer) GetItems(ctx context.Context, req *GetItemsRequest) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}

func RegisterDcmApiDatabaseServiceServer(s *grpc.Server, srv DcmApiDatabaseServiceServer) {
	s.RegisterService(&_DcmApiDatabaseService_serviceDesc, srv)
}

func _DcmApiDatabaseService_InsertItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcmApiDatabaseServiceServer).InsertItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DcmApiDatabaseService/InsertItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcmApiDatabaseServiceServer).InsertItem(ctx, req.(*InsertItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DcmApiDatabaseService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcmApiDatabaseServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DcmApiDatabaseService/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcmApiDatabaseServiceServer).GetItems(ctx, req.(*GetItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DcmApiDatabaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DcmApiDatabaseService",
	HandlerType: (*DcmApiDatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertItem",
			Handler:    _DcmApiDatabaseService_InsertItem_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _DcmApiDatabaseService_GetItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
