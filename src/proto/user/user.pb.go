// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserListRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListRequest) Reset()         { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()    {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListRequest.Unmarshal(m, b)
}
func (m *GetUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListRequest.Marshal(b, m, deterministic)
}
func (m *GetUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListRequest.Merge(m, src)
}
func (m *GetUserListRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserListRequest.Size(m)
}
func (m *GetUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListRequest proto.InternalMessageInfo

func (m *GetUserListRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserListRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserCourseResponse struct {
	IsSuccess            bool      `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Message              string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Userinfo             *UserInfo `protobuf:"bytes,3,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetUserCourseResponse) Reset()         { *m = GetUserCourseResponse{} }
func (m *GetUserCourseResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserCourseResponse) ProtoMessage()    {}
func (*GetUserCourseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserCourseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCourseResponse.Unmarshal(m, b)
}
func (m *GetUserCourseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCourseResponse.Marshal(b, m, deterministic)
}
func (m *GetUserCourseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCourseResponse.Merge(m, src)
}
func (m *GetUserCourseResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserCourseResponse.Size(m)
}
func (m *GetUserCourseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCourseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCourseResponse proto.InternalMessageInfo

func (m *GetUserCourseResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *GetUserCourseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetUserCourseResponse) GetUserinfo() *UserInfo {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

type UserInfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserListRequest)(nil), "user.GetUserListRequest")
	proto.RegisterType((*GetUserCourseResponse)(nil), "user.GetUserCourseResponse")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x25, 0xf1, 0x6b, 0x77, 0x16, 0x7b, 0x18, 0x10, 0x42, 0x45, 0x58, 0xf6, 0xb4, 0x78, 0xe8,
	0xa1, 0x82, 0x67, 0x41, 0x41, 0x04, 0x4f, 0x59, 0x3c, 0x4b, 0xdd, 0x4e, 0x25, 0x87, 0x26, 0x35,
	0x93, 0x78, 0xf2, 0xc7, 0x4b, 0xb2, 0x6d, 0x31, 0x78, 0xcb, 0xbc, 0x47, 0xde, 0x17, 0x40, 0x64,
	0xf2, 0x8b, 0x9d, 0x77, 0xc1, 0xe1, 0x69, 0x7a, 0x77, 0x0f, 0x80, 0xcf, 0x14, 0xde, 0x98, 0xfc,
	0xab, 0xe1, 0xa0, 0xe9, 0x2b, 0x12, 0x07, 0x9c, 0x43, 0x95, 0x58, 0xbb, 0xda, 0x92, 0x12, 0xad,
	0xe8, 0x6b, 0x7d, 0xbc, 0x71, 0x06, 0xd2, 0xac, 0x95, 0x6c, 0x45, 0x7f, 0xa6, 0xa5, 0x59, 0x77,
	0x3f, 0x70, 0xb5, 0x57, 0x78, 0x74, 0xd1, 0x33, 0x69, 0xe2, 0x9d, 0xb3, 0x4c, 0x78, 0x03, 0x60,
	0xf8, 0x9d, 0xe3, 0x38, 0x12, 0x73, 0x96, 0xa9, 0x74, 0x6d, 0x78, 0x98, 0x00, 0x54, 0x70, 0xb1,
	0x25, 0xe6, 0xd5, 0x27, 0x65, 0xb1, 0x5a, 0x1f, 0x4e, 0xbc, 0x9d, 0xdc, 0x8d, 0xdd, 0x38, 0x75,
	0xd2, 0x8a, 0xbe, 0x59, 0xce, 0x16, 0x39, 0x78, 0x32, 0x79, 0xb1, 0x1b, 0xa7, 0x8f, 0x7c, 0x77,
	0x0f, 0xd5, 0x01, 0xdd, 0x27, 0x4b, 0x46, 0x97, 0x29, 0x59, 0xd1, 0x42, 0x96, 0x2d, 0x96, 0x03,
	0x34, 0xe9, 0xdf, 0x40, 0xfe, 0xdb, 0x8c, 0x84, 0x4f, 0xd0, 0xfc, 0x99, 0x01, 0xd5, 0xe4, 0xf7,
	0x7f, 0x99, 0xf9, 0x75, 0xc1, 0x94, 0x8d, 0x3f, 0xce, 0xf3, 0xb2, 0x77, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0x19, 0xf5, 0x29, 0x67, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserCourseResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserCourseResponse, error) {
	out := new(GetUserCourseResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUserList(context.Context, *GetUserListRequest) (*GetUserCourseResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _UserService_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
