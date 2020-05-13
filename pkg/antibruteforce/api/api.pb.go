// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

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

type NetAddr struct {
	NetAddr              string   `protobuf:"bytes,1,opt,name=netAddr,proto3" json:"netAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetAddr) Reset()         { *m = NetAddr{} }
func (m *NetAddr) String() string { return proto.CompactTextString(m) }
func (*NetAddr) ProtoMessage()    {}
func (*NetAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *NetAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetAddr.Unmarshal(m, b)
}
func (m *NetAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetAddr.Marshal(b, m, deterministic)
}
func (m *NetAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetAddr.Merge(m, src)
}
func (m *NetAddr) XXX_Size() int {
	return xxx_messageInfo_NetAddr.Size(m)
}
func (m *NetAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_NetAddr.DiscardUnknown(m)
}

var xxx_messageInfo_NetAddr proto.InternalMessageInfo

func (m *NetAddr) GetNetAddr() string {
	if m != nil {
		return m.NetAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*NetAddr)(nil), "NetAddr")
}

func init() {
	proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784)
}

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0x92, 0x4a,
	0xca, 0x5c, 0xec, 0x7e, 0xa9, 0x25, 0x8e, 0x29, 0x29, 0x45, 0x42, 0x12, 0x5c, 0xec, 0x79, 0x10,
	0xa6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x6b, 0x74, 0x8b, 0x91, 0x8b, 0xcf, 0x31,
	0xaf, 0x24, 0x33, 0xa9, 0xa8, 0xb4, 0x24, 0x35, 0x2d, 0xbf, 0x28, 0x39, 0x55, 0x48, 0x9f, 0x8b,
	0xdb, 0x31, 0x25, 0xc5, 0x29, 0x27, 0x31, 0x39, 0xdb, 0x2f, 0xb5, 0x44, 0x88, 0x43, 0x0f, 0x6a,
	0x8a, 0x94, 0x98, 0x1e, 0xc4, 0x3a, 0x3d, 0x98, 0x75, 0x7a, 0xae, 0x20, 0xeb, 0x84, 0x8c, 0xb8,
	0xf8, 0x82, 0x52, 0x73, 0xf3, 0xcb, 0x52, 0x49, 0xd0, 0x03, 0xb1, 0x24, 0x3c, 0x23, 0xb3, 0x24,
	0x95, 0x44, 0x4b, 0x88, 0xd7, 0x93, 0xc4, 0x06, 0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x55, 0xf5, 0xf0, 0x30, 0x36, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AntibruteforceClient is the client API for Antibruteforce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AntibruteforceClient interface {
	AddBlackNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveBlackNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error)
	AddWhiteNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveWhiteNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error)
}

type antibruteforceClient struct {
	cc grpc.ClientConnInterface
}

func NewAntibruteforceClient(cc grpc.ClientConnInterface) AntibruteforceClient {
	return &antibruteforceClient{cc}
}

func (c *antibruteforceClient) AddBlackNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Antibruteforce/AddBlackNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antibruteforceClient) RemoveBlackNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Antibruteforce/RemoveBlackNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antibruteforceClient) AddWhiteNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Antibruteforce/AddWhiteNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antibruteforceClient) RemoveWhiteNet(ctx context.Context, in *NetAddr, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Antibruteforce/RemoveWhiteNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AntibruteforceServer is the server API for Antibruteforce service.
type AntibruteforceServer interface {
	AddBlackNet(context.Context, *NetAddr) (*empty.Empty, error)
	RemoveBlackNet(context.Context, *NetAddr) (*empty.Empty, error)
	AddWhiteNet(context.Context, *NetAddr) (*empty.Empty, error)
	RemoveWhiteNet(context.Context, *NetAddr) (*empty.Empty, error)
}

// UnimplementedAntibruteforceServer can be embedded to have forward compatible implementations.
type UnimplementedAntibruteforceServer struct {
}

func (*UnimplementedAntibruteforceServer) AddBlackNet(ctx context.Context, req *NetAddr) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlackNet not implemented")
}
func (*UnimplementedAntibruteforceServer) RemoveBlackNet(ctx context.Context, req *NetAddr) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlackNet not implemented")
}
func (*UnimplementedAntibruteforceServer) AddWhiteNet(ctx context.Context, req *NetAddr) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWhiteNet not implemented")
}
func (*UnimplementedAntibruteforceServer) RemoveWhiteNet(ctx context.Context, req *NetAddr) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWhiteNet not implemented")
}

func RegisterAntibruteforceServer(s *grpc.Server, srv AntibruteforceServer) {
	s.RegisterService(&_Antibruteforce_serviceDesc, srv)
}

func _Antibruteforce_AddBlackNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntibruteforceServer).AddBlackNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Antibruteforce/AddBlackNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntibruteforceServer).AddBlackNet(ctx, req.(*NetAddr))
	}
	return interceptor(ctx, in, info, handler)
}

func _Antibruteforce_RemoveBlackNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntibruteforceServer).RemoveBlackNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Antibruteforce/RemoveBlackNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntibruteforceServer).RemoveBlackNet(ctx, req.(*NetAddr))
	}
	return interceptor(ctx, in, info, handler)
}

func _Antibruteforce_AddWhiteNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntibruteforceServer).AddWhiteNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Antibruteforce/AddWhiteNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntibruteforceServer).AddWhiteNet(ctx, req.(*NetAddr))
	}
	return interceptor(ctx, in, info, handler)
}

func _Antibruteforce_RemoveWhiteNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntibruteforceServer).RemoveWhiteNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Antibruteforce/RemoveWhiteNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntibruteforceServer).RemoveWhiteNet(ctx, req.(*NetAddr))
	}
	return interceptor(ctx, in, info, handler)
}

var _Antibruteforce_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Antibruteforce",
	HandlerType: (*AntibruteforceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlackNet",
			Handler:    _Antibruteforce_AddBlackNet_Handler,
		},
		{
			MethodName: "RemoveBlackNet",
			Handler:    _Antibruteforce_RemoveBlackNet_Handler,
		},
		{
			MethodName: "AddWhiteNet",
			Handler:    _Antibruteforce_AddWhiteNet_Handler,
		},
		{
			MethodName: "RemoveWhiteNet",
			Handler:    _Antibruteforce_RemoveWhiteNet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
