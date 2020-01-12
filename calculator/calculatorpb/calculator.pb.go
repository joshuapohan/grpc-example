// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CalculatorInput struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorInput) Reset()         { *m = CalculatorInput{} }
func (m *CalculatorInput) String() string { return proto.CompactTextString(m) }
func (*CalculatorInput) ProtoMessage()    {}
func (*CalculatorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *CalculatorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorInput.Unmarshal(m, b)
}
func (m *CalculatorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorInput.Marshal(b, m, deterministic)
}
func (m *CalculatorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorInput.Merge(m, src)
}
func (m *CalculatorInput) XXX_Size() int {
	return xxx_messageInfo_CalculatorInput.Size(m)
}
func (m *CalculatorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorInput.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorInput proto.InternalMessageInfo

func (m *CalculatorInput) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *CalculatorInput) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type CalculatorRequest struct {
	CalculatorInput      *CalculatorInput `protobuf:"bytes,1,opt,name=calculator_input,json=calculatorInput,proto3" json:"calculator_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetCalculatorInput() *CalculatorInput {
	if m != nil {
		return m.CalculatorInput
	}
	return nil
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AverageRequest struct {
	Input                int32    `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRequest) Reset()         { *m = AverageRequest{} }
func (m *AverageRequest) String() string { return proto.CompactTextString(m) }
func (*AverageRequest) ProtoMessage()    {}
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *AverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRequest.Unmarshal(m, b)
}
func (m *AverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRequest.Marshal(b, m, deterministic)
}
func (m *AverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRequest.Merge(m, src)
}
func (m *AverageRequest) XXX_Size() int {
	return xxx_messageInfo_AverageRequest.Size(m)
}
func (m *AverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRequest proto.InternalMessageInfo

func (m *AverageRequest) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type AverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (m *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(m, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MaximumRequest struct {
	Input                int32    `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaximumRequest) Reset()         { *m = MaximumRequest{} }
func (m *MaximumRequest) String() string { return proto.CompactTextString(m) }
func (*MaximumRequest) ProtoMessage()    {}
func (*MaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *MaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximumRequest.Unmarshal(m, b)
}
func (m *MaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximumRequest.Marshal(b, m, deterministic)
}
func (m *MaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximumRequest.Merge(m, src)
}
func (m *MaximumRequest) XXX_Size() int {
	return xxx_messageInfo_MaximumRequest.Size(m)
}
func (m *MaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MaximumRequest proto.InternalMessageInfo

func (m *MaximumRequest) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type MaximumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaximumResponse) Reset()         { *m = MaximumResponse{} }
func (m *MaximumResponse) String() string { return proto.CompactTextString(m) }
func (*MaximumResponse) ProtoMessage()    {}
func (*MaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *MaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximumResponse.Unmarshal(m, b)
}
func (m *MaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximumResponse.Marshal(b, m, deterministic)
}
func (m *MaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximumResponse.Merge(m, src)
}
func (m *MaximumResponse) XXX_Size() int {
	return xxx_messageInfo_MaximumResponse.Size(m)
}
func (m *MaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaximumResponse proto.InternalMessageInfo

func (m *MaximumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type SquareRootRequest struct {
	Input                int32    `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type SquareRootResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{8}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculatorInput)(nil), "calculator.CalculatorInput")
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
	proto.RegisterType((*AverageRequest)(nil), "calculator.AverageRequest")
	proto.RegisterType((*AverageResponse)(nil), "calculator.AverageResponse")
	proto.RegisterType((*MaximumRequest)(nil), "calculator.MaximumRequest")
	proto.RegisterType((*MaximumResponse)(nil), "calculator.MaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdd, 0x4a, 0x02, 0x41,
	0x18, 0x86, 0x1b, 0x43, 0x85, 0x2f, 0xf1, 0xe7, 0x23, 0x42, 0x94, 0x22, 0xe6, 0x20, 0x14, 0xcc,
	0xc2, 0xae, 0xc0, 0x02, 0x29, 0xa2, 0x93, 0xf5, 0xac, 0x0e, 0x62, 0x76, 0x1d, 0x42, 0x58, 0x9d,
	0x75, 0x76, 0x46, 0xba, 0xb5, 0xee, 0x2e, 0xd6, 0xfd, 0x99, 0x19, 0xdd, 0xf2, 0x6c, 0xdf, 0x99,
	0x6f, 0x9f, 0x77, 0xe6, 0x81, 0x81, 0x41, 0xc0, 0xc2, 0x40, 0x87, 0x4c, 0x09, 0x79, 0x67, 0x3e,
	0x23, 0xdf, 0x0a, 0xe3, 0x48, 0x0a, 0x25, 0x10, 0xcc, 0x0a, 0xbd, 0x85, 0xd6, 0x53, 0x91, 0x5e,
	0xd6, 0x91, 0x56, 0xd8, 0x00, 0xc2, 0xba, 0xe4, 0x9a, 0x0c, 0xaa, 0x1e, 0x61, 0x49, 0xf2, 0xbb,
	0x95, 0x34, 0xf9, 0xf4, 0x03, 0x3a, 0x66, 0xdc, 0xe3, 0x1b, 0xcd, 0x63, 0x85, 0x33, 0x68, 0x1b,
	0xe2, 0xe7, 0x32, 0x81, 0xec, 0xfe, 0x3f, 0x9b, 0xf4, 0xc7, 0x56, 0xf9, 0x5e, 0x8f, 0xd7, 0x0a,
	0xdc, 0x05, 0x3a, 0x02, 0xb4, 0xe1, 0x71, 0x24, 0xd6, 0x31, 0xc7, 0x0b, 0xa8, 0x49, 0x1e, 0xeb,
	0x50, 0x65, 0x67, 0xca, 0x12, 0xbd, 0x81, 0xe6, 0x74, 0xcb, 0x25, 0xfb, 0xe2, 0xf9, 0x39, 0xce,
	0xa1, 0x6a, 0xca, 0xab, 0x5e, 0x1a, 0xe8, 0x10, 0x5a, 0xc5, 0x5c, 0x29, 0xb2, 0x62, 0x23, 0xdf,
	0xd8, 0xf7, 0x72, 0xa5, 0x57, 0x47, 0x91, 0xc5, 0xdc, 0x91, 0x53, 0x0e, 0xa1, 0x33, 0xdf, 0x68,
	0x26, 0xb9, 0x27, 0x84, 0xfa, 0x9f, 0x3a, 0x02, 0xb4, 0x47, 0x4b, 0xc1, 0x24, 0x07, 0x4f, 0x7e,
	0x2a, 0xd0, 0xce, 0x6d, 0xf1, 0x39, 0x97, 0xdb, 0x65, 0xc0, 0x71, 0x06, 0xa7, 0xd3, 0xc5, 0x02,
	0x2f, 0xcb, 0xb5, 0x67, 0xf5, 0xbd, 0xab, 0xbf, 0xb6, 0xd3, 0x4a, 0x7a, 0x82, 0x33, 0xa8, 0x67,
	0xce, 0xb0, 0x67, 0x0f, 0xbb, 0xc2, 0x7b, 0xfd, 0xd2, 0xbd, 0x9c, 0x32, 0x20, 0xf8, 0x0c, 0xf5,
	0x4c, 0x94, 0xcb, 0x71, 0x2d, 0xbb, 0x9c, 0x3d, 0xb3, 0x09, 0xe7, 0x9e, 0xe0, 0x2b, 0x80, 0x91,
	0xe3, 0x5e, 0xf0, 0xc0, 0xaf, 0x7b, 0xc1, 0x43, 0xa7, 0x8f, 0xcd, 0xf7, 0x86, 0xfd, 0x42, 0xfc,
	0xda, 0xee, 0x5d, 0x3c, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x01, 0xe1, 0x43, 0x78, 0x43, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Add(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AverageClient, error)
	Maximum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_MaximumClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculateServiceClient(cc *grpc.ClientConn) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Add(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[0], "/calculator.CalculateService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceAverageClient{stream}
	return x, nil
}

type CalculateService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculateServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculateServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) Maximum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_MaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[1], "/calculator.CalculateService/Maximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceMaximumClient{stream}
	return x, nil
}

type CalculateService_MaximumClient interface {
	Send(*MaximumRequest) error
	Recv() (*MaximumResponse, error)
	grpc.ClientStream
}

type calculateServiceMaximumClient struct {
	grpc.ClientStream
}

func (x *calculateServiceMaximumClient) Send(m *MaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceMaximumClient) Recv() (*MaximumResponse, error) {
	m := new(MaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	Add(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	Average(CalculateService_AverageServer) error
	Maximum(CalculateService_MaximumServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) Add(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculateServiceServer) Average(srv CalculateService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedCalculateServiceServer) Maximum(srv CalculateService_MaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method Maximum not implemented")
}
func (*UnimplementedCalculateServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Add(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).Average(&calculateServiceAverageServer{stream})
}

type CalculateService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculateServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculateServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_Maximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).Maximum(&calculateServiceMaximumServer{stream})
}

type CalculateService_MaximumServer interface {
	Send(*MaximumResponse) error
	Recv() (*MaximumRequest, error)
	grpc.ServerStream
}

type calculateServiceMaximumServer struct {
	grpc.ServerStream
}

func (x *calculateServiceMaximumServer) Send(m *MaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceMaximumServer) Recv() (*MaximumRequest, error) {
	m := new(MaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculateService_Add_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Average",
			Handler:       _CalculateService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Maximum",
			Handler:       _CalculateService_Maximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
