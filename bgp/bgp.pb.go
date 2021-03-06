// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openconfig/gnoi/bgp/bgp.proto

/*
Package gnoi_bgp is a generated protocol buffer package.

It is generated from these files:
	github.com/openconfig/gnoi/bgp/bgp.proto

It has these top-level messages:
	ClearBGPNeighborRequest
	ClearBGPNeighborResponse
*/
package gnoi_bgp

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

type ClearBGPNeighborRequest_Mode int32

const (
	ClearBGPNeighborRequest_SOFT   ClearBGPNeighborRequest_Mode = 0
	ClearBGPNeighborRequest_SOFTIN ClearBGPNeighborRequest_Mode = 1
	ClearBGPNeighborRequest_HARD   ClearBGPNeighborRequest_Mode = 2
)

var ClearBGPNeighborRequest_Mode_name = map[int32]string{
	0: "SOFT",
	1: "SOFTIN",
	2: "HARD",
}
var ClearBGPNeighborRequest_Mode_value = map[string]int32{
	"SOFT":   0,
	"SOFTIN": 1,
	"HARD":   2,
}

func (x ClearBGPNeighborRequest_Mode) String() string {
	return proto.EnumName(ClearBGPNeighborRequest_Mode_name, int32(x))
}
func (ClearBGPNeighborRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type ClearBGPNeighborRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Routing instance containing the neighbor. Defaults to the global routing
	// table.
	RoutingInstance string                       `protobuf:"bytes,2,opt,name=routing_instance,json=routingInstance" json:"routing_instance,omitempty"`
	Mode            ClearBGPNeighborRequest_Mode `protobuf:"varint,3,opt,name=mode,enum=gnoi.bgp.ClearBGPNeighborRequest_Mode" json:"mode,omitempty"`
}

func (m *ClearBGPNeighborRequest) Reset()                    { *m = ClearBGPNeighborRequest{} }
func (m *ClearBGPNeighborRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearBGPNeighborRequest) ProtoMessage()               {}
func (*ClearBGPNeighborRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClearBGPNeighborRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClearBGPNeighborRequest) GetRoutingInstance() string {
	if m != nil {
		return m.RoutingInstance
	}
	return ""
}

func (m *ClearBGPNeighborRequest) GetMode() ClearBGPNeighborRequest_Mode {
	if m != nil {
		return m.Mode
	}
	return ClearBGPNeighborRequest_SOFT
}

type ClearBGPNeighborResponse struct {
}

func (m *ClearBGPNeighborResponse) Reset()                    { *m = ClearBGPNeighborResponse{} }
func (m *ClearBGPNeighborResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearBGPNeighborResponse) ProtoMessage()               {}
func (*ClearBGPNeighborResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ClearBGPNeighborRequest)(nil), "gnoi.bgp.ClearBGPNeighborRequest")
	proto.RegisterType((*ClearBGPNeighborResponse)(nil), "gnoi.bgp.ClearBGPNeighborResponse")
	proto.RegisterEnum("gnoi.bgp.ClearBGPNeighborRequest_Mode", ClearBGPNeighborRequest_Mode_name, ClearBGPNeighborRequest_Mode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BGP service

type BGPClient interface {
	// ClearBGPNeighbor clears a BGP session.
	ClearBGPNeighbor(ctx context.Context, in *ClearBGPNeighborRequest, opts ...grpc.CallOption) (*ClearBGPNeighborResponse, error)
}

type bGPClient struct {
	cc *grpc.ClientConn
}

func NewBGPClient(cc *grpc.ClientConn) BGPClient {
	return &bGPClient{cc}
}

func (c *bGPClient) ClearBGPNeighbor(ctx context.Context, in *ClearBGPNeighborRequest, opts ...grpc.CallOption) (*ClearBGPNeighborResponse, error) {
	out := new(ClearBGPNeighborResponse)
	err := grpc.Invoke(ctx, "/gnoi.bgp.BGP/ClearBGPNeighbor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BGP service

type BGPServer interface {
	// ClearBGPNeighbor clears a BGP session.
	ClearBGPNeighbor(context.Context, *ClearBGPNeighborRequest) (*ClearBGPNeighborResponse, error)
}

func RegisterBGPServer(s *grpc.Server, srv BGPServer) {
	s.RegisterService(&_BGP_serviceDesc, srv)
}

func _BGP_ClearBGPNeighbor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearBGPNeighborRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGPServer).ClearBGPNeighbor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.bgp.BGP/ClearBGPNeighbor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGPServer).ClearBGPNeighbor(ctx, req.(*ClearBGPNeighborRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BGP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.bgp.BGP",
	HandlerType: (*BGPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClearBGPNeighbor",
			Handler:    _BGP_ClearBGPNeighbor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openconfig/gnoi/bgp/bgp.proto",
}

func init() { proto.RegisterFile("github.com/openconfig/gnoi/bgp/bgp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x9b, 0x36, 0x2a, 0xe5, 0x06, 0x88, 0xbc, 0x10, 0x75, 0x2a, 0x19, 0xaa, 0xb0, 0x38,
	0x52, 0xd9, 0xd8, 0x28, 0x88, 0xd2, 0x81, 0x52, 0x05, 0x36, 0x06, 0x14, 0x27, 0x87, 0x6b, 0x89,
	0xfa, 0x8c, 0xed, 0xfc, 0x41, 0x7e, 0x19, 0x72, 0x68, 0x17, 0x10, 0x74, 0x38, 0xe9, 0xde, 0xa7,
	0x77, 0xf7, 0x4e, 0x07, 0xb9, 0x54, 0x7e, 0xd3, 0x0a, 0x5e, 0xd3, 0xb6, 0x20, 0x83, 0xba, 0x26,
	0xfd, 0xa6, 0x64, 0x21, 0x35, 0xa9, 0x42, 0x48, 0x13, 0x8a, 0x1b, 0x4b, 0x9e, 0xd8, 0x28, 0x30,
	0x2e, 0xa4, 0xc9, 0x3e, 0x23, 0x38, 0xbb, 0x79, 0xc7, 0xca, 0xce, 0x17, 0xeb, 0x15, 0x2a, 0xb9,
	0x11, 0x64, 0x4b, 0xfc, 0x68, 0xd1, 0x79, 0x96, 0xc2, 0x51, 0xd5, 0x34, 0x16, 0x9d, 0x4b, 0xa3,
	0x49, 0x94, 0x1f, 0x97, 0x7b, 0xc9, 0x2e, 0x20, 0xb1, 0xd4, 0x7a, 0xa5, 0xe5, 0xab, 0xd2, 0xce,
	0x57, 0xba, 0xc6, 0xb4, 0xdf, 0x59, 0x4e, 0x77, 0x7c, 0xb9, 0xc3, 0xec, 0x0a, 0xe2, 0x2d, 0x35,
	0x98, 0x0e, 0x26, 0x51, 0x7e, 0x32, 0x9b, 0xf2, 0x7d, 0x32, 0xff, 0x23, 0x95, 0x3f, 0x50, 0x83,
	0x65, 0x37, 0x93, 0x4d, 0x21, 0x0e, 0x8a, 0x8d, 0x20, 0x7e, 0x7a, 0xbc, 0x7b, 0x4e, 0x7a, 0x0c,
	0x60, 0x18, 0xba, 0xe5, 0x2a, 0x89, 0x02, 0xbd, 0xbf, 0x2e, 0x6f, 0x93, 0x7e, 0x36, 0x86, 0xf4,
	0xf7, 0x36, 0x67, 0x48, 0x3b, 0x9c, 0x09, 0x18, 0xcc, 0x17, 0x6b, 0xf6, 0x02, 0xc9, 0x4f, 0x0b,
	0x3b, 0x3f, 0x78, 0xcc, 0x38, 0xfb, 0xcf, 0xf2, 0x9d, 0x90, 0xf5, 0xc4, 0xb0, 0xfb, 0xea, 0xe5,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xe1, 0xbb, 0xd9, 0x81, 0x01, 0x00, 0x00,
}
