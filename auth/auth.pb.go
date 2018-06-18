// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	User
*/
package auth

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
const _ = proto.ProtoPackageIsVersion1

// The request message containing the username and password.
type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the JWT token.
type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The user message containing the user.
type User struct {
	Email        string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	PasswordHash string `protobuf:"bytes,3,opt,name=passwordHash" json:"passwordHash,omitempty"`
	IsAdmin      bool   `protobuf:"varint,4,opt,name=isAdmin" json:"isAdmin,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*User)(nil), "auth.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Auth service

type AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/auth.Auth/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServer).Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xdc, 0xb8, 0x78, 0x7c, 0xf2,
	0xd3, 0x33, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0x8b,
	0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90,
	0x5c, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13, 0x44, 0x0e, 0xc6, 0x57, 0x52,
	0xe5, 0xe2, 0x85, 0x9a, 0x53, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc2, 0xc5, 0x5a, 0x92,
	0x9f, 0x9d, 0x9a, 0x07, 0x35, 0x05, 0xc2, 0x51, 0x2a, 0xe3, 0x62, 0x09, 0x05, 0x1a, 0x07, 0x92,
	0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x81, 0xc9, 0x82, 0x39, 0x28, 0x96, 0x33, 0xa1, 0x59, 0xae, 0xc4,
	0xc5, 0x03, 0xb3, 0xcc, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x19, 0x2c, 0x8f, 0x22, 0x26, 0x24, 0xc1,
	0xc5, 0x9e, 0x59, 0xec, 0x98, 0x92, 0x9b, 0x99, 0x27, 0xc1, 0x02, 0x94, 0xe6, 0x08, 0x82, 0x71,
	0x8d, 0x2c, 0xb8, 0x58, 0x1c, 0x81, 0xde, 0x15, 0x32, 0xe0, 0x62, 0x05, 0x3b, 0x53, 0x48, 0x48,
	0x0f, 0x1c, 0x14, 0xc8, 0x7e, 0x97, 0x12, 0x46, 0x11, 0x83, 0xf8, 0x23, 0x89, 0x0d, 0x1c, 0x5a,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xef, 0x10, 0x7a, 0x3b, 0x01, 0x00, 0x00,
}
