// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfsre/cfsre/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cfsre/cfsre/tx.proto", fileDescriptor_b2c79866b2a64b61) }

var fileDescriptor_b2c79866b2a64b61 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4e, 0x2b, 0x2e,
	0x4a, 0xd5, 0x87, 0x90, 0x25, 0x15, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xa2, 0x65, 0xa9,
	0x29, 0x19, 0x89, 0x65, 0x95, 0x89, 0xc5, 0x7a, 0x60, 0x19, 0x08, 0x69, 0xc4, 0xca, 0xc5, 0xec,
	0x5b, 0x9c, 0xee, 0xe4, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x9a,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x70, 0x23, 0xf4, 0x93, 0xd3,
	0x74, 0x41, 0xa6, 0x57, 0xc0, 0x6c, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xdb, 0x64, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd8, 0x60, 0xa0, 0x81, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vedhavyas.cfsre.cfsre.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cfsre/cfsre/tx.proto",
}
