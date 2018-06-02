// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package models

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

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_40e9213e78b15951, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type User struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_40e9213e78b15951, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserWithPassword struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserWithPassword) Reset()         { *m = UserWithPassword{} }
func (m *UserWithPassword) String() string { return proto.CompactTextString(m) }
func (*UserWithPassword) ProtoMessage()    {}
func (*UserWithPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_40e9213e78b15951, []int{2}
}
func (m *UserWithPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserWithPassword.Unmarshal(m, b)
}
func (m *UserWithPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserWithPassword.Marshal(b, m, deterministic)
}
func (dst *UserWithPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserWithPassword.Merge(dst, src)
}
func (m *UserWithPassword) XXX_Size() int {
	return xxx_messageInfo_UserWithPassword.Size(m)
}
func (m *UserWithPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_UserWithPassword.DiscardUnknown(m)
}

var xxx_messageInfo_UserWithPassword proto.InternalMessageInfo

func (m *UserWithPassword) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserWithPassword) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "proto.Session")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserWithPassword)(nil), "proto.UserWithPassword")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Check(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error)
	Login(ctx context.Context, in *UserWithPassword, opts ...grpc.CallOption) (*Session, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Check(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.Auth/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *UserWithPassword, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/proto.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Check(context.Context, *Session) (*User, error)
	Login(context.Context, *UserWithPassword) (*Session, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Check(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserWithPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*UserWithPassword))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Auth_Check_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_40e9213e78b15951) }

var fileDescriptor_auth_40e9213e78b15951 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xea, 0x5c, 0xec, 0xc1, 0xa9,
	0xc5, 0xc5, 0x99, 0xf9, 0x79, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0x10, 0xa6, 0x67, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x42, 0x40, 0x49, 0x86, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48, 0x48,
	0x84, 0x8b, 0x35, 0x35, 0x37, 0x31, 0x33, 0x07, 0xaa, 0x02, 0xc2, 0x51, 0x72, 0xe1, 0x12, 0x00,
	0xc9, 0x86, 0x67, 0x96, 0x64, 0x04, 0x24, 0x16, 0x17, 0x97, 0xe7, 0x17, 0xa5, 0x60, 0x57, 0x29,
	0x24, 0xc5, 0xc5, 0x51, 0x00, 0x55, 0x21, 0xc1, 0x04, 0x96, 0x80, 0xf3, 0x8d, 0x92, 0xb8, 0x58,
	0x1c, 0x4b, 0x4b, 0x32, 0x84, 0xd4, 0xb8, 0x58, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x85, 0xf8, 0x20,
	0x8e, 0xd5, 0x83, 0x3a, 0x51, 0x8a, 0x1b, 0xca, 0x07, 0xd9, 0xa5, 0xc4, 0x20, 0x64, 0xc4, 0xc5,
	0xea, 0x93, 0x9f, 0x9e, 0x99, 0x27, 0x24, 0x8e, 0x24, 0x8e, 0xec, 0x06, 0x29, 0x34, 0x03, 0x94,
	0x18, 0x92, 0xd8, 0xc0, 0x02, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x25, 0xc9, 0x0f,
	0x0c, 0x01, 0x00, 0x00,
}
