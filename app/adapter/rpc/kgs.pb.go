// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/adapter/rpc/kgs.proto

package rpc

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PopulateKeysRequest struct {
	KeySize              uint32   `protobuf:"varint,1,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopulateKeysRequest) Reset()         { *m = PopulateKeysRequest{} }
func (m *PopulateKeysRequest) String() string { return proto.CompactTextString(m) }
func (*PopulateKeysRequest) ProtoMessage()    {}
func (*PopulateKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb3e7329511ae677, []int{0}
}

func (m *PopulateKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopulateKeysRequest.Unmarshal(m, b)
}
func (m *PopulateKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopulateKeysRequest.Marshal(b, m, deterministic)
}
func (m *PopulateKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopulateKeysRequest.Merge(m, src)
}
func (m *PopulateKeysRequest) XXX_Size() int {
	return xxx_messageInfo_PopulateKeysRequest.Size(m)
}
func (m *PopulateKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PopulateKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PopulateKeysRequest proto.InternalMessageInfo

func (m *PopulateKeysRequest) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

type PopulateKeysResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopulateKeysResponse) Reset()         { *m = PopulateKeysResponse{} }
func (m *PopulateKeysResponse) String() string { return proto.CompactTextString(m) }
func (*PopulateKeysResponse) ProtoMessage()    {}
func (*PopulateKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb3e7329511ae677, []int{1}
}

func (m *PopulateKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopulateKeysResponse.Unmarshal(m, b)
}
func (m *PopulateKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopulateKeysResponse.Marshal(b, m, deterministic)
}
func (m *PopulateKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopulateKeysResponse.Merge(m, src)
}
func (m *PopulateKeysResponse) XXX_Size() int {
	return xxx_messageInfo_PopulateKeysResponse.Size(m)
}
func (m *PopulateKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PopulateKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PopulateKeysResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PopulateKeysRequest)(nil), "rpc.PopulateKeysRequest")
	proto.RegisterType((*PopulateKeysResponse)(nil), "rpc.PopulateKeysResponse")
}

func init() { proto.RegisterFile("app/adapter/rpc/kgs.proto", fileDescriptor_eb3e7329511ae677) }

var fileDescriptor_eb3e7329511ae677 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0xcd, 0x0a, 0x82, 0x40,
	0x14, 0x85, 0x91, 0xc0, 0x62, 0xa8, 0xcd, 0x14, 0xa2, 0xb6, 0x09, 0x57, 0xad, 0x66, 0xa2, 0x5e,
	0xa0, 0x4d, 0x04, 0xb9, 0x09, 0x7b, 0x80, 0x18, 0xed, 0x26, 0xa2, 0x39, 0xb7, 0xf9, 0x59, 0x8c,
	0x4f, 0x1f, 0x29, 0x41, 0x41, 0xcb, 0x73, 0xf8, 0x0e, 0xe7, 0x23, 0x91, 0x40, 0xe4, 0xe2, 0x26,
	0xd0, 0x80, 0xe2, 0x0a, 0x0b, 0x5e, 0x97, 0x9a, 0xa1, 0x92, 0x46, 0xd2, 0x91, 0xc2, 0x22, 0x5e,
	0x96, 0x52, 0x96, 0x0d, 0xf0, 0xbe, 0xca, 0xed, 0x9d, 0xc3, 0x03, 0x8d, 0x1b, 0x88, 0x64, 0x43,
	0xe6, 0x67, 0x89, 0xb6, 0x11, 0x06, 0x52, 0x70, 0x3a, 0x83, 0xa7, 0x05, 0x6d, 0x68, 0x44, 0x26,
	0x35, 0xb8, 0xab, 0xae, 0x3a, 0x08, 0xbd, 0x95, 0xb7, 0x9e, 0x65, 0xe3, 0x1a, 0xdc, 0xa5, 0xea,
	0x20, 0x09, 0xc8, 0xe2, 0x77, 0xa1, 0x51, 0xb6, 0x1a, 0xb6, 0x27, 0xe2, 0xa7, 0xe0, 0x8e, 0xd0,
	0xd2, 0x3d, 0x99, 0x7e, 0x13, 0x34, 0x64, 0x0a, 0x0b, 0xf6, 0xe7, 0x26, 0x0e, 0xd8, 0xe0, 0xc6,
	0x3e, 0x6e, 0xec, 0xf0, 0x76, 0xcb, 0xfd, 0x3e, 0xef, 0x5e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7a,
	0xbc, 0xa0, 0xfb, 0xdb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyGenClient is the client API for KeyGen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyGenClient interface {
	PopulateKeys(ctx context.Context, in *PopulateKeysRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type keyGenClient struct {
	cc *grpc.ClientConn
}

func NewKeyGenClient(cc *grpc.ClientConn) KeyGenClient {
	return &keyGenClient{cc}
}

func (c *keyGenClient) PopulateKeys(ctx context.Context, in *PopulateKeysRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpc.KeyGen/PopulateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyGenServer is the server API for KeyGen service.
type KeyGenServer interface {
	PopulateKeys(context.Context, *PopulateKeysRequest) (*empty.Empty, error)
}

// UnimplementedKeyGenServer can be embedded to have forward compatible implementations.
type UnimplementedKeyGenServer struct {
}

func (*UnimplementedKeyGenServer) PopulateKeys(ctx context.Context, req *PopulateKeysRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopulateKeys not implemented")
}

func RegisterKeyGenServer(s *grpc.Server, srv KeyGenServer) {
	s.RegisterService(&_KeyGen_serviceDesc, srv)
}

func _KeyGen_PopulateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopulateKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyGenServer).PopulateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KeyGen/PopulateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyGenServer).PopulateKeys(ctx, req.(*PopulateKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyGen_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.KeyGen",
	HandlerType: (*KeyGenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PopulateKeys",
			Handler:    _KeyGen_PopulateKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/adapter/rpc/kgs.proto",
}
