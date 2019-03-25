// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microservice/microservicepb/microservice.proto

package microservicepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Microservice struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Microservice) Reset()         { *m = Microservice{} }
func (m *Microservice) String() string { return proto.CompactTextString(m) }
func (*Microservice) ProtoMessage()    {}
func (*Microservice) Descriptor() ([]byte, []int) {
	return fileDescriptor_microservice_b78ff35aeb513d86, []int{0}
}
func (m *Microservice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Microservice.Unmarshal(m, b)
}
func (m *Microservice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Microservice.Marshal(b, m, deterministic)
}
func (dst *Microservice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Microservice.Merge(dst, src)
}
func (m *Microservice) XXX_Size() int {
	return xxx_messageInfo_Microservice.Size(m)
}
func (m *Microservice) XXX_DiscardUnknown() {
	xxx_messageInfo_Microservice.DiscardUnknown(m)
}

var xxx_messageInfo_Microservice proto.InternalMessageInfo

func (m *Microservice) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Microservice) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type MicroserviceRequest struct {
	Microservice         *Microservice `protobuf:"bytes,1,opt,name=microservice,proto3" json:"microservice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MicroserviceRequest) Reset()         { *m = MicroserviceRequest{} }
func (m *MicroserviceRequest) String() string { return proto.CompactTextString(m) }
func (*MicroserviceRequest) ProtoMessage()    {}
func (*MicroserviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_microservice_b78ff35aeb513d86, []int{1}
}
func (m *MicroserviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MicroserviceRequest.Unmarshal(m, b)
}
func (m *MicroserviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MicroserviceRequest.Marshal(b, m, deterministic)
}
func (dst *MicroserviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicroserviceRequest.Merge(dst, src)
}
func (m *MicroserviceRequest) XXX_Size() int {
	return xxx_messageInfo_MicroserviceRequest.Size(m)
}
func (m *MicroserviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MicroserviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MicroserviceRequest proto.InternalMessageInfo

func (m *MicroserviceRequest) GetMicroservice() *Microservice {
	if m != nil {
		return m.Microservice
	}
	return nil
}

type MicroserviceResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MicroserviceResponse) Reset()         { *m = MicroserviceResponse{} }
func (m *MicroserviceResponse) String() string { return proto.CompactTextString(m) }
func (*MicroserviceResponse) ProtoMessage()    {}
func (*MicroserviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_microservice_b78ff35aeb513d86, []int{2}
}
func (m *MicroserviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MicroserviceResponse.Unmarshal(m, b)
}
func (m *MicroserviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MicroserviceResponse.Marshal(b, m, deterministic)
}
func (dst *MicroserviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicroserviceResponse.Merge(dst, src)
}
func (m *MicroserviceResponse) XXX_Size() int {
	return xxx_messageInfo_MicroserviceResponse.Size(m)
}
func (m *MicroserviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MicroserviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MicroserviceResponse proto.InternalMessageInfo

func (m *MicroserviceResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Microservice)(nil), "microservice.Microservice")
	proto.RegisterType((*MicroserviceRequest)(nil), "microservice.MicroserviceRequest")
	proto.RegisterType((*MicroserviceResponse)(nil), "microservice.MicroserviceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MicroserviceServiceClient is the client API for MicroserviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MicroserviceServiceClient interface {
	// Clent Streaming
	Microservice(ctx context.Context, opts ...grpc.CallOption) (MicroserviceService_MicroserviceClient, error)
}

type microserviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewMicroserviceServiceClient(cc *grpc.ClientConn) MicroserviceServiceClient {
	return &microserviceServiceClient{cc}
}

func (c *microserviceServiceClient) Microservice(ctx context.Context, opts ...grpc.CallOption) (MicroserviceService_MicroserviceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MicroserviceService_serviceDesc.Streams[0], "/microservice.MicroserviceService/Microservice", opts...)
	if err != nil {
		return nil, err
	}
	x := &microserviceServiceMicroserviceClient{stream}
	return x, nil
}

type MicroserviceService_MicroserviceClient interface {
	Send(*MicroserviceRequest) error
	CloseAndRecv() (*MicroserviceResponse, error)
	grpc.ClientStream
}

type microserviceServiceMicroserviceClient struct {
	grpc.ClientStream
}

func (x *microserviceServiceMicroserviceClient) Send(m *MicroserviceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *microserviceServiceMicroserviceClient) CloseAndRecv() (*MicroserviceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MicroserviceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MicroserviceServiceServer is the server API for MicroserviceService service.
type MicroserviceServiceServer interface {
	// Clent Streaming
	Microservice(MicroserviceService_MicroserviceServer) error
}

func RegisterMicroserviceServiceServer(s *grpc.Server, srv MicroserviceServiceServer) {
	s.RegisterService(&_MicroserviceService_serviceDesc, srv)
}

func _MicroserviceService_Microservice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MicroserviceServiceServer).Microservice(&microserviceServiceMicroserviceServer{stream})
}

type MicroserviceService_MicroserviceServer interface {
	SendAndClose(*MicroserviceResponse) error
	Recv() (*MicroserviceRequest, error)
	grpc.ServerStream
}

type microserviceServiceMicroserviceServer struct {
	grpc.ServerStream
}

func (x *microserviceServiceMicroserviceServer) SendAndClose(m *MicroserviceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *microserviceServiceMicroserviceServer) Recv() (*MicroserviceRequest, error) {
	m := new(MicroserviceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MicroserviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "microservice.MicroserviceService",
	HandlerType: (*MicroserviceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Microservice",
			Handler:       _MicroserviceService_Microservice_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "microservice/microservicepb/microservice.proto",
}

func init() {
	proto.RegisterFile("microservice/microservicepb/microservice.proto", fileDescriptor_microservice_b78ff35aeb513d86)
}

var fileDescriptor_microservice_b78ff35aeb513d86 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0xcb, 0xcd, 0x4c, 0x2e,
	0xca, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x47, 0xe6, 0x14, 0x24, 0xa1, 0x70, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x90, 0xc5, 0x94, 0xbc, 0xb8, 0x78, 0x7c, 0x91, 0xf8,
	0x42, 0xb2, 0x5c, 0x5c, 0x69, 0x99, 0x45, 0xc5, 0x25, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x9c, 0x60, 0x11, 0x3f, 0xa0, 0x80, 0x90, 0x34, 0x17, 0x67, 0x4e,
	0x22, 0x4c, 0x96, 0x09, 0x2c, 0xcb, 0x01, 0x12, 0x00, 0x49, 0x2a, 0x85, 0x72, 0x09, 0x23, 0x9b,
	0x15, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x64, 0xc7, 0x85, 0x62, 0x25, 0xd8, 0x50, 0x6e,
	0x23, 0x29, 0x14, 0x77, 0xeb, 0xa1, 0x68, 0x44, 0x75, 0xa2, 0x1e, 0x97, 0x08, 0xaa, 0xb1, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x40, 0x76, 0x69, 0x4e, 0x09, 0xd4, 0x99,
	0x50, 0x9e, 0x51, 0x01, 0xaa, 0x33, 0x82, 0xa1, 0x3e, 0x8b, 0x44, 0xf3, 0xa9, 0x22, 0x1e, 0x07,
	0x40, 0x5c, 0x2e, 0xa5, 0x84, 0x4f, 0x09, 0xc4, 0x15, 0x4a, 0x0c, 0x1a, 0x8c, 0x4e, 0x02, 0x51,
	0x7c, 0xa8, 0xe1, 0x9e, 0xc4, 0x06, 0x0e, 0x6b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x73, 0xb5, 0x6c, 0x9d, 0x01, 0x00, 0x00,
}
