// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oauth.proto

package oauth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// User registred in "xmas"
type User struct {
	// Uuua claim to put in jwt
	Uuaa string `protobuf:"bytes,1,opt,name=uuaa,proto3" json:"uuaa,omitempty"`
	// IVU claim to put in jwt
	Ivuser               string   `protobuf:"bytes,2,opt,name=ivuser,proto3" json:"ivuser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{0}
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

func (m *User) GetUuaa() string {
	if m != nil {
		return m.Uuaa
	}
	return ""
}

func (m *User) GetIvuser() string {
	if m != nil {
		return m.Ivuser
	}
	return ""
}

// Request data to create new JWT
type CreateRequest struct {
	// Task entity to add
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Response that contains data for created JWT
type CreateResponse struct {
	// Specification of type of JWT(In this case will be Bearer)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Acces_token is the token value
	AccessToken          string   `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "oauth.User")
	proto.RegisterType((*CreateRequest)(nil), "oauth.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "oauth.CreateResponse")
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor_7ce0b12f599e9f07) }

var fileDescriptor_7ce0b12f599e9f07 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0xb0, 0xdd, 0x0a, 0xaf, 0x8a, 0xc0, 0x2a, 0xb0, 0x8a, 0x2a, 0x61, 0xe5, 0x84,
	0x56, 0x34, 0x5e, 0xc2, 0x9e, 0x96, 0x0b, 0x05, 0xc1, 0x05, 0x24, 0x50, 0x0a, 0x87, 0x1e, 0x4d,
	0x3a, 0x4d, 0x5c, 0xb2, 0x1e, 0xe3, 0xb1, 0x77, 0xe9, 0x95, 0x47, 0x80, 0x1b, 0xaf, 0xc5, 0x85,
	0x03, 0x47, 0x1e, 0x04, 0xc5, 0x9b, 0x22, 0xa0, 0x27, 0x8f, 0xff, 0xff, 0x9f, 0xf9, 0x46, 0x1a,
	0x36, 0x41, 0x15, 0x7c, 0x5b, 0x58, 0x87, 0x1e, 0xf9, 0x4e, 0xfc, 0x64, 0x07, 0x0d, 0x62, 0xd3,
	0x81, 0x54, 0x56, 0x4b, 0x65, 0x0c, 0x7a, 0xe5, 0x35, 0x1a, 0xda, 0x86, 0xb2, 0x07, 0xf1, 0xa9,
	0x0f, 0x1b, 0x30, 0x87, 0xb4, 0x51, 0x4d, 0x03, 0x4e, 0xa2, 0x8d, 0x89, 0xab, 0xe9, 0xbc, 0x64,
	0xa3, 0x77, 0x04, 0x8e, 0x73, 0x36, 0x0a, 0x41, 0xa9, 0x69, 0x22, 0x92, 0xfb, 0xd7, 0xab, 0x58,
	0xf3, 0x3b, 0x6c, 0xac, 0xd7, 0x81, 0xc0, 0x4d, 0xd3, 0xa8, 0x0e, 0xbf, 0x7c, 0xce, 0xf6, 0x9e,
	0x39, 0x50, 0x1e, 0x2a, 0xf8, 0x18, 0x80, 0x3c, 0xbf, 0xc7, 0x46, 0x7f, 0x62, 0x93, 0x72, 0x52,
	0x6c, 0x77, 0xee, 0xe7, 0x56, 0xd1, 0xc8, 0x5f, 0xb0, 0x1b, 0x97, 0x1d, 0x64, 0xd1, 0x10, 0xf4,
	0x3c, 0x7f, 0x61, 0xe1, 0x92, 0xd7, 0xd7, 0x5c, 0xb0, 0xc9, 0x51, 0x5d, 0x03, 0xd1, 0x5b, 0xfc,
	0x00, 0x66, 0x80, 0xfe, 0x2d, 0x95, 0xe7, 0x6c, 0xef, 0x75, 0x3f, 0xbb, 0x3c, 0x06, 0xb7, 0xd6,
	0x35, 0xf0, 0x13, 0x36, 0xde, 0x0e, 0xe6, 0xfb, 0x03, 0xf5, 0x9f, 0xcd, 0xb2, 0xdb, 0xff, 0xa9,
	0x5b, 0x7a, 0x9e, 0x7f, 0xfe, 0xfe, 0xeb, 0x6b, 0x7a, 0x90, 0xdf, 0x95, 0xd1, 0x96, 0x0f, 0x8b,
	0x79, 0x31, 0x97, 0xb1, 0xf4, 0x3d, 0x68, 0x99, 0xcc, 0x9e, 0xfe, 0x4c, 0xbe, 0x1c, 0xfd, 0x48,
	0x78, 0xcb, 0x6e, 0xbd, 0x41, 0xe7, 0x83, 0x09, 0x24, 0x06, 0x2a, 0xe5, 0xaf, 0xd8, 0xf4, 0xe5,
	0xf3, 0x93, 0x63, 0x71, 0x86, 0x4e, 0xac, 0x2e, 0x84, 0x81, 0x8d, 0xa0, 0xc1, 0xe3, 0xfb, 0xad,
	0xf7, 0x96, 0x96, 0x52, 0xda, 0xa1, 0xad, 0xa8, 0x71, 0x95, 0x65, 0x67, 0x9d, 0x5a, 0xeb, 0x40,
	0x45, 0xa7, 0x6a, 0xe5, 0x03, 0x85, 0x27, 0xcd, 0x4a, 0xe9, 0xae, 0xf7, 0xca, 0x9d, 0x88, 0xcf,
	0x77, 0x63, 0x87, 0x56, 0xb3, 0x34, 0x49, 0xcb, 0x9b, 0xca, 0xda, 0x4e, 0xd7, 0xf1, 0x56, 0xf2,
	0x9c, 0xd0, 0x2c, 0xaf, 0x28, 0xd5, 0x63, 0x76, 0x6d, 0x31, 0x5f, 0xf0, 0x05, 0x9b, 0x55, 0xe0,
	0x83, 0x33, 0x70, 0x2a, 0x36, 0x2d, 0x18, 0xe1, 0x5b, 0x10, 0x0e, 0x08, 0x83, 0xab, 0x41, 0x9c,
	0x22, 0x90, 0x30, 0xe8, 0x05, 0x7c, 0xd2, 0xe4, 0x0b, 0x3e, 0x66, 0xa3, 0x6f, 0x69, 0xb2, 0xfb,
	0x7e, 0x1c, 0xef, 0xff, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x13, 0x73, 0xd1, 0x61,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Oauth2ServiceClient is the client API for Oauth2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Oauth2ServiceClient interface {
	// Create new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type oauth2ServiceClient struct {
	cc *grpc.ClientConn
}

func NewOauth2ServiceClient(cc *grpc.ClientConn) Oauth2ServiceClient {
	return &oauth2ServiceClient{cc}
}

func (c *oauth2ServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/oauth.Oauth2Service/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Oauth2ServiceServer is the server API for Oauth2Service service.
type Oauth2ServiceServer interface {
	// Create new todo task
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedOauth2ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOauth2ServiceServer struct {
}

func (*UnimplementedOauth2ServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterOauth2ServiceServer(s *grpc.Server, srv Oauth2ServiceServer) {
	s.RegisterService(&_Oauth2Service_serviceDesc, srv)
}

func _Oauth2Service_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2ServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth2Service/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2ServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Oauth2Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.Oauth2Service",
	HandlerType: (*Oauth2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Oauth2Service_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}