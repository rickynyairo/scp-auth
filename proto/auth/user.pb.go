// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/user.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08e205afd2d1b12, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08e205afd2d1b12, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08e205afd2d1b12, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08e205afd2d1b12, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_f08e205afd2d1b12, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Response)(nil), "auth.Response")
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*Error)(nil), "auth.Error")
}

func init() { proto.RegisterFile("proto/auth/user.proto", fileDescriptor_f08e205afd2d1b12) }

var fileDescriptor_f08e205afd2d1b12 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0x6e, 0xf3, 0xab, 0xed, 0x94, 0xf6, 0x30, 0xbc, 0x07, 0x4b, 0x0f, 0x8f, 0xbc, 0x14, 0x44,
	0x11, 0x5a, 0xa8, 0x67, 0x0f, 0x45, 0xa4, 0xf7, 0xa0, 0xe2, 0x75, 0x6d, 0x06, 0x1a, 0x4c, 0xb3,
	0xe9, 0xee, 0xa6, 0xe2, 0x3f, 0xe9, 0xdf, 0x24, 0x3b, 0x49, 0x25, 0x14, 0x41, 0x4f, 0x99, 0xef,
	0x07, 0x93, 0x6f, 0x3e, 0x16, 0xfe, 0x56, 0x5a, 0x59, 0xb5, 0x94, 0xb5, 0xdd, 0x2d, 0x6b, 0x43,
	0x7a, 0xc1, 0x18, 0x03, 0x47, 0x24, 0x47, 0x08, 0x1e, 0x0d, 0x69, 0x9c, 0x82, 0x97, 0x67, 0xa2,
	0x1f, 0xf7, 0x2f, 0x47, 0xa9, 0x97, 0x67, 0x88, 0x10, 0x94, 0x72, 0x4f, 0xc2, 0x63, 0x86, 0x67,
	0x14, 0x30, 0xd8, 0xaa, 0x7d, 0x25, 0xcb, 0x77, 0xe1, 0x33, 0x7d, 0x82, 0xf8, 0x07, 0x42, 0xda,
	0xcb, 0xbc, 0x10, 0x01, 0xf3, 0x0d, 0xc0, 0x19, 0x0c, 0x2b, 0x69, 0xcc, 0x9b, 0xd2, 0x99, 0x08,
	0x59, 0xf8, 0xc2, 0xc9, 0x08, 0x06, 0x29, 0x1d, 0x6a, 0x32, 0x36, 0x39, 0xc0, 0x30, 0x25, 0x53,
	0xa9, 0xd2, 0x10, 0xfe, 0x83, 0xc0, 0x45, 0xe4, 0x20, 0xe3, 0x15, 0x2c, 0x5c, 0xc6, 0x85, 0x0b,
	0x98, 0x32, 0x8f, 0x31, 0x84, 0xee, 0x6b, 0x84, 0x17, 0xfb, 0x67, 0x86, 0x46, 0xc0, 0x39, 0x44,
	0xa4, 0xb5, 0xd2, 0x46, 0xf8, 0x6c, 0x19, 0x37, 0x96, 0x7b, 0xc7, 0xa5, 0xad, 0x94, 0x3c, 0x43,
	0xf8, 0xa0, 0x5e, 0xa9, 0x74, 0xc1, 0xad, 0x1b, 0xda, 0xcb, 0x1b, 0xe0, 0xd8, 0xa3, 0x2c, 0xf2,
	0x8c, 0xaf, 0x1f, 0xa6, 0x0d, 0xf8, 0xdd, 0xe6, 0x5b, 0x08, 0x99, 0x70, 0x05, 0x6e, 0x55, 0x46,
	0xbc, 0x38, 0x4c, 0x79, 0xc6, 0x18, 0xc6, 0x19, 0x99, 0xad, 0xce, 0x2b, 0x9b, 0xab, 0xb2, 0xed,
	0xb6, 0x4b, 0xad, 0x3e, 0xfa, 0x10, 0xac, 0x6b, 0xbb, 0xc3, 0x0b, 0x88, 0xee, 0x34, 0x49, 0x4b,
	0xd8, 0xb9, 0x71, 0x36, 0x6d, 0xe6, 0x53, 0x5d, 0x49, 0x0f, 0xe7, 0xe0, 0x6f, 0xc8, 0xfe, 0x60,
	0xba, 0x82, 0x68, 0x43, 0x76, 0x5d, 0x14, 0x38, 0x39, 0x69, 0x5c, 0xfd, 0x37, 0xd6, 0xff, 0xed,
	0xff, 0xbb, 0x0b, 0xdb, 0x43, 0xb9, 0xb1, 0xa4, 0x87, 0xd7, 0x30, 0x79, 0x72, 0x85, 0x48, 0x4b,
	0x4d, 0x89, 0x5d, 0xfd, 0xcc, 0xfc, 0x12, 0xf1, 0x63, 0xbb, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x65, 0x66, 0x1a, 0x85, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Create(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *Auth) Get(ctx context.Context, in *User, out *Response) error {
	return h.AuthHandler.Get(ctx, in, out)
}

func (h *Auth) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.GetAll(ctx, in, out)
}

func (h *Auth) Auth(ctx context.Context, in *User, out *Token) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *Auth) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}
