// Code generated by protoc-gen-go. DO NOT EDIT.
// source: played.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	played.proto

It has these top-level messages:
	SendPlayedRequest
	GameEntry
	SendPlayedResponse
	GetPlayedRequest
	GetPlayedResponse
*/
package pb

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

type SendPlayedRequest struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Game string `protobuf:"bytes,2,opt,name=game" json:"game,omitempty"`
}

func (m *SendPlayedRequest) Reset()                    { *m = SendPlayedRequest{} }
func (m *SendPlayedRequest) String() string            { return proto.CompactTextString(m) }
func (*SendPlayedRequest) ProtoMessage()               {}
func (*SendPlayedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendPlayedRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SendPlayedRequest) GetGame() string {
	if m != nil {
		return m.Game
	}
	return ""
}

type GameEntry struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Dur  int64  `protobuf:"varint,2,opt,name=dur" json:"dur,omitempty"`
}

func (m *GameEntry) Reset()                    { *m = GameEntry{} }
func (m *GameEntry) String() string            { return proto.CompactTextString(m) }
func (*GameEntry) ProtoMessage()               {}
func (*GameEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GameEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameEntry) GetDur() int64 {
	if m != nil {
		return m.Dur
	}
	return 0
}

type SendPlayedResponse struct {
}

func (m *SendPlayedResponse) Reset()                    { *m = SendPlayedResponse{} }
func (m *SendPlayedResponse) String() string            { return proto.CompactTextString(m) }
func (*SendPlayedResponse) ProtoMessage()               {}
func (*SendPlayedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetPlayedRequest struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetPlayedRequest) Reset()                    { *m = GetPlayedRequest{} }
func (m *GetPlayedRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPlayedRequest) ProtoMessage()               {}
func (*GetPlayedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetPlayedRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetPlayedResponse struct {
}

func (m *GetPlayedResponse) Reset()                    { *m = GetPlayedResponse{} }
func (m *GetPlayedResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPlayedResponse) ProtoMessage()               {}
func (*GetPlayedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*SendPlayedRequest)(nil), "played.SendPlayedRequest")
	proto.RegisterType((*GameEntry)(nil), "played.GameEntry")
	proto.RegisterType((*SendPlayedResponse)(nil), "played.SendPlayedResponse")
	proto.RegisterType((*GetPlayedRequest)(nil), "played.GetPlayedRequest")
	proto.RegisterType((*GetPlayedResponse)(nil), "played.GetPlayedResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Played service

type PlayedClient interface {
	SendPlayed(ctx context.Context, opts ...grpc.CallOption) (Played_SendPlayedClient, error)
	GetPlayed(ctx context.Context, in *GetPlayedRequest, opts ...grpc.CallOption) (*GetPlayedResponse, error)
}

type playedClient struct {
	cc *grpc.ClientConn
}

func NewPlayedClient(cc *grpc.ClientConn) PlayedClient {
	return &playedClient{cc}
}

func (c *playedClient) SendPlayed(ctx context.Context, opts ...grpc.CallOption) (Played_SendPlayedClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Played_serviceDesc.Streams[0], c.cc, "/played.Played/SendPlayed", opts...)
	if err != nil {
		return nil, err
	}
	x := &playedSendPlayedClient{stream}
	return x, nil
}

type Played_SendPlayedClient interface {
	Send(*SendPlayedRequest) error
	CloseAndRecv() (*SendPlayedResponse, error)
	grpc.ClientStream
}

type playedSendPlayedClient struct {
	grpc.ClientStream
}

func (x *playedSendPlayedClient) Send(m *SendPlayedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *playedSendPlayedClient) CloseAndRecv() (*SendPlayedResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendPlayedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *playedClient) GetPlayed(ctx context.Context, in *GetPlayedRequest, opts ...grpc.CallOption) (*GetPlayedResponse, error) {
	out := new(GetPlayedResponse)
	err := grpc.Invoke(ctx, "/played.Played/GetPlayed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Played service

type PlayedServer interface {
	SendPlayed(Played_SendPlayedServer) error
	GetPlayed(context.Context, *GetPlayedRequest) (*GetPlayedResponse, error)
}

func RegisterPlayedServer(s *grpc.Server, srv PlayedServer) {
	s.RegisterService(&_Played_serviceDesc, srv)
}

func _Played_SendPlayed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlayedServer).SendPlayed(&playedSendPlayedServer{stream})
}

type Played_SendPlayedServer interface {
	SendAndClose(*SendPlayedResponse) error
	Recv() (*SendPlayedRequest, error)
	grpc.ServerStream
}

type playedSendPlayedServer struct {
	grpc.ServerStream
}

func (x *playedSendPlayedServer) SendAndClose(m *SendPlayedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *playedSendPlayedServer) Recv() (*SendPlayedRequest, error) {
	m := new(SendPlayedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Played_GetPlayed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayedServer).GetPlayed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/played.Played/GetPlayed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayedServer).GetPlayed(ctx, req.(*GetPlayedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Played_serviceDesc = grpc.ServiceDesc{
	ServiceName: "played.Played",
	HandlerType: (*PlayedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayed",
			Handler:    _Played_GetPlayed_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendPlayed",
			Handler:       _Played_SendPlayed_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "played.proto",
}

func init() { proto.RegisterFile("played.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x49, 0xac,
	0x4c, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xac, 0xb9, 0x04,
	0x83, 0x53, 0xf3, 0x52, 0x02, 0xc0, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21,
	0x2e, 0x96, 0xd2, 0xe2, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x1b, 0x24,
	0x96, 0x9e, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x11, 0x03, 0xb1, 0x95, 0x0c, 0xb9, 0x38, 0xdd, 0x13,
	0x73, 0x53, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x41, 0x0a, 0xf2, 0x40, 0x0a, 0xa0, 0x9a, 0x40, 0x6c,
	0x21, 0x01, 0x2e, 0xe6, 0x94, 0xd2, 0x22, 0xb0, 0x1e, 0xe6, 0x20, 0x10, 0x53, 0x49, 0x84, 0x4b,
	0x08, 0xd9, 0xbe, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0x35, 0x2e, 0x01, 0xf7, 0xd4, 0x12,
	0x82, 0x8e, 0x50, 0x12, 0xe6, 0x12, 0x44, 0x52, 0x07, 0xd1, 0x6c, 0x34, 0x99, 0x91, 0x8b, 0x0d,
	0x22, 0x24, 0xe4, 0xce, 0xc5, 0x85, 0x30, 0x5d, 0x48, 0x52, 0x0f, 0xea, 0x65, 0x0c, 0x1f, 0x4a,
	0x49, 0x61, 0x93, 0x82, 0x3a, 0x86, 0x41, 0x83, 0x51, 0xc8, 0x81, 0x8b, 0x13, 0x6e, 0x91, 0x90,
	0x04, 0x4c, 0x31, 0xba, 0x1b, 0xa5, 0x24, 0xb1, 0xc8, 0x40, 0x4c, 0x71, 0x62, 0x89, 0x62, 0x2a,
	0x48, 0x4a, 0x62, 0x03, 0x87, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x84, 0xe3, 0x01, 0x5b,
	0x7d, 0x01, 0x00, 0x00,
}
