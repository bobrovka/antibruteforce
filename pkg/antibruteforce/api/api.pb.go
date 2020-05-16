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

type CheckRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *CheckRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CheckRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type Status struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*NetAddr)(nil), "NetAddr")
	proto.RegisterType((*CheckRequest)(nil), "CheckRequest")
	proto.RegisterType((*Status)(nil), "Status")
}

func init() {
	proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784)
}

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x55, 0x83, 0xda, 0x84, 0x83, 0x66, 0xb0, 0x2a, 0x14, 0x85, 0x01, 0x14, 0x16, 0x26, 0x47,
	0x2a, 0xbf, 0x20, 0x20, 0x98, 0x50, 0x85, 0xc2, 0xc0, 0x9c, 0x8f, 0x6b, 0x6a, 0x25, 0x8d, 0x8d,
	0x7d, 0x01, 0xf1, 0xfb, 0xf8, 0x63, 0xa8, 0x76, 0x8b, 0xba, 0x74, 0xc8, 0xe6, 0xe7, 0x77, 0xef,
	0xee, 0xde, 0x3d, 0x98, 0x17, 0x4a, 0xa4, 0x85, 0x12, 0x5c, 0x69, 0x49, 0x32, 0xbe, 0x6e, 0xa4,
	0x6c, 0x3a, 0x4c, 0x2d, 0x2a, 0x87, 0x75, 0x8a, 0x5b, 0x45, 0x3f, 0x8e, 0x4c, 0xee, 0xc0, 0x5f,
	0x21, 0x65, 0x75, 0xad, 0x59, 0x04, 0x7e, 0xef, 0x9e, 0xd1, 0xe4, 0x76, 0x72, 0x7f, 0x9e, 0x1f,
	0x60, 0xf2, 0x06, 0x97, 0x4f, 0x1b, 0xac, 0xda, 0x1c, 0x3f, 0x07, 0x34, 0xc4, 0x16, 0x30, 0xed,
	0x64, 0x23, 0xfa, 0x7d, 0x9d, 0x03, 0x2c, 0x86, 0x40, 0x15, 0xc6, 0x7c, 0x4b, 0x5d, 0x47, 0x9e,
	0x25, 0xfe, 0x31, 0x0b, 0xc1, 0x13, 0x2a, 0x3a, 0xb3, 0xbf, 0x9e, 0x50, 0x49, 0x04, 0xb3, 0x77,
	0x2a, 0x68, 0x30, 0x3b, 0x46, 0xb6, 0xb6, 0x51, 0x90, 0x7b, 0xb2, 0x5d, 0xfe, 0x7a, 0x10, 0x66,
	0x3d, 0x89, 0x52, 0x0f, 0x84, 0x6b, 0xa9, 0x2b, 0x64, 0x29, 0x5c, 0x64, 0x75, 0xfd, 0xd8, 0x15,
	0x55, 0xbb, 0x42, 0x62, 0x01, 0xdf, 0x6f, 0x1c, 0x5f, 0x71, 0x67, 0x8d, 0x1f, 0xac, 0xf1, 0xe7,
	0x9d, 0x35, 0xb6, 0x84, 0x30, 0xc7, 0xad, 0xfc, 0xc2, 0x11, 0x1a, 0x37, 0xe4, 0x63, 0x23, 0x08,
	0x47, 0x0e, 0x19, 0xa1, 0x79, 0x81, 0x45, 0x8e, 0x06, 0xc9, 0xee, 0x65, 0x75, 0xaf, 0xc2, 0x90,
	0x61, 0x27, 0xea, 0x4f, 0xf6, 0xb9, 0x81, 0xa9, 0x0d, 0x84, 0xcd, 0xf9, 0x71, 0x30, 0xb1, 0xcf,
	0xdd, 0x55, 0xcb, 0x99, 0x15, 0x3c, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0x2b, 0x27, 0xff,
	0x0b, 0x02, 0x00, 0x00,
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
	ResetBlackWhiteLists(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Status, error)
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

func (c *antibruteforceClient) ResetBlackWhiteLists(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Antibruteforce/ResetBlackWhiteLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antibruteforceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Antibruteforce/Check", in, out, opts...)
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
	ResetBlackWhiteLists(context.Context, *empty.Empty) (*empty.Empty, error)
	Check(context.Context, *CheckRequest) (*Status, error)
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
func (*UnimplementedAntibruteforceServer) ResetBlackWhiteLists(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetBlackWhiteLists not implemented")
}
func (*UnimplementedAntibruteforceServer) Check(ctx context.Context, req *CheckRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
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

func _Antibruteforce_ResetBlackWhiteLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntibruteforceServer).ResetBlackWhiteLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Antibruteforce/ResetBlackWhiteLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntibruteforceServer).ResetBlackWhiteLists(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Antibruteforce_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntibruteforceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Antibruteforce/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntibruteforceServer).Check(ctx, req.(*CheckRequest))
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
		{
			MethodName: "ResetBlackWhiteLists",
			Handler:    _Antibruteforce_ResetBlackWhiteLists_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Antibruteforce_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
