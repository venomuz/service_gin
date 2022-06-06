// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: post.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type AllPost struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllPost) Reset()         { *m = AllPost{} }
func (m *AllPost) String() string { return proto.CompactTextString(m) }
func (*AllPost) ProtoMessage()    {}
func (*AllPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{0}
}
func (m *AllPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllPost.Merge(m, src)
}
func (m *AllPost) XXX_Size() int {
	return m.Size()
}
func (m *AllPost) XXX_DiscardUnknown() {
	xxx_messageInfo_AllPost.DiscardUnknown(m)
}

var xxx_messageInfo_AllPost proto.InternalMessageInfo

func (m *AllPost) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Medias               []*Media `protobuf:"bytes,5,rep,name=medias,proto3" json:"medias"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{1}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Post.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return m.Size()
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Post) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Post) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Post) GetMedias() []*Media {
	if m != nil {
		return m.Medias
	}
	return nil
}

type Media struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	PostId               string   `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	Link                 string   `protobuf:"bytes,4,opt,name=link,proto3" json:"link"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}
func (*Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{2}
}
func (m *Media) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Media.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Media.Merge(m, src)
}
func (m *Media) XXX_Size() int {
	return m.Size()
}
func (m *Media) XXX_DiscardUnknown() {
	xxx_messageInfo_Media.DiscardUnknown(m)
}

var xxx_messageInfo_Media proto.InternalMessageInfo

func (m *Media) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Media) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *Media) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Media) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type GetIdFromUser struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIdFromUser) Reset()         { *m = GetIdFromUser{} }
func (m *GetIdFromUser) String() string { return proto.CompactTextString(m) }
func (*GetIdFromUser) ProtoMessage()    {}
func (*GetIdFromUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{3}
}
func (m *GetIdFromUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetIdFromUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetIdFromUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetIdFromUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIdFromUser.Merge(m, src)
}
func (m *GetIdFromUser) XXX_Size() int {
	return m.Size()
}
func (m *GetIdFromUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIdFromUser.DiscardUnknown(m)
}

var xxx_messageInfo_GetIdFromUser proto.InternalMessageInfo

func (m *GetIdFromUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type OkBOOL struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OkBOOL) Reset()         { *m = OkBOOL{} }
func (m *OkBOOL) String() string { return proto.CompactTextString(m) }
func (*OkBOOL) ProtoMessage()    {}
func (*OkBOOL) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{4}
}
func (m *OkBOOL) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OkBOOL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OkBOOL.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OkBOOL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkBOOL.Merge(m, src)
}
func (m *OkBOOL) XXX_Size() int {
	return m.Size()
}
func (m *OkBOOL) XXX_DiscardUnknown() {
	xxx_messageInfo_OkBOOL.DiscardUnknown(m)
}

var xxx_messageInfo_OkBOOL proto.InternalMessageInfo

func (m *OkBOOL) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*AllPost)(nil), "user.AllPost")
	proto.RegisterType((*Post)(nil), "user.Post")
	proto.RegisterType((*Media)(nil), "user.Media")
	proto.RegisterType((*GetIdFromUser)(nil), "user.GetIdFromUser")
	proto.RegisterType((*OkBOOL)(nil), "user.OkBOOL")
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x7f, 0xdb, 0xa6, 0xe9, 0xcf, 0x89, 0xad, 0xb2, 0x82, 0x0d, 0x3d, 0xc4, 0x10, 0x41,
	0x0a, 0x42, 0x0e, 0x2d, 0x3e, 0x80, 0xb5, 0x58, 0x02, 0x4a, 0x25, 0x22, 0x78, 0x93, 0xd8, 0xcc,
	0x61, 0x69, 0xda, 0x84, 0xdd, 0xad, 0xd0, 0x57, 0xf0, 0x09, 0x7c, 0x24, 0x8f, 0xfa, 0x06, 0x52,
	0x5f, 0x44, 0xf6, 0xcf, 0xa1, 0xf5, 0xcf, 0x69, 0x67, 0xbe, 0xdf, 0x99, 0x9d, 0x4f, 0x26, 0x0b,
	0x50, 0x95, 0x42, 0xc6, 0x15, 0x2f, 0x65, 0x49, 0x9d, 0xa5, 0x40, 0x1e, 0x9d, 0x42, 0xf3, 0xbc,
	0x28, 0x6e, 0x4a, 0x21, 0x69, 0x08, 0x0d, 0x75, 0x0a, 0x9f, 0x84, 0xf5, 0x9e, 0xd7, 0x87, 0x58,
	0x15, 0xc4, 0x4a, 0x4a, 0x8d, 0x11, 0x3d, 0x13, 0x70, 0x74, 0x69, 0x1b, 0x6a, 0x2c, 0xf7, 0x49,
	0x48, 0x7a, 0x3b, 0x69, 0x8d, 0xe5, 0x94, 0x82, 0xb3, 0xc8, 0xe6, 0xe8, 0xd7, 0xb4, 0xa2, 0x63,
	0x1a, 0x82, 0x97, 0xa3, 0x98, 0x72, 0x56, 0x49, 0x56, 0x2e, 0xfc, 0xba, 0xb6, 0x36, 0x25, 0xda,
	0x81, 0xa6, 0x1a, 0xf1, 0xc0, 0x72, 0xdf, 0xd1, 0xae, 0xab, 0xd2, 0x24, 0xa7, 0xc7, 0xe0, 0xce,
	0x31, 0x67, 0x99, 0xf0, 0x1b, 0x1a, 0xc5, 0x33, 0x28, 0xd7, 0x4a, 0x4b, 0xad, 0x15, 0xdd, 0x43,
	0x43, 0x0b, 0x3f, 0x60, 0x3a, 0xd0, 0x54, 0x9f, 0xa9, 0xae, 0x35, 0x3c, 0xae, 0x4a, 0x13, 0x4d,
	0x29, 0x57, 0x15, 0x5a, 0x14, 0x1d, 0x2b, 0xad, 0x60, 0x8b, 0x99, 0x05, 0xd0, 0x71, 0x74, 0x04,
	0xad, 0x31, 0xca, 0x24, 0xbf, 0xe4, 0xe5, 0xfc, 0x4e, 0x20, 0xff, 0x3e, 0x21, 0x0a, 0xc1, 0x9d,
	0xcc, 0x86, 0x93, 0xc9, 0x15, 0x3d, 0x04, 0x57, 0xc8, 0x4c, 0x2e, 0x85, 0x76, 0xff, 0xa7, 0x36,
	0xeb, 0xbf, 0x13, 0xf0, 0xd4, 0xa6, 0x6e, 0x91, 0x3f, 0xb1, 0x29, 0xd2, 0x13, 0x00, 0x95, 0x5e,
	0x70, 0xcc, 0x24, 0xd2, 0x8d, 0xd5, 0x76, 0x77, 0x4d, 0x6c, 0xef, 0x8b, 0x4d, 0xdb, 0x18, 0xe5,
	0x70, 0x95, 0x8c, 0xe8, 0x81, 0x31, 0xb7, 0x68, 0xba, 0x1b, 0xdd, 0x74, 0x00, 0x6d, 0x75, 0x8e,
	0xb0, 0x40, 0x89, 0x7f, 0xb7, 0x6c, 0x0f, 0x39, 0x83, 0x3d, 0x3b, 0xc4, 0xfe, 0x7a, 0xf1, 0x7b,
	0x57, 0xcb, 0x88, 0xb6, 0x68, 0xb8, 0xff, 0xba, 0x0e, 0xc8, 0xdb, 0x3a, 0x20, 0x1f, 0xeb, 0x80,
	0xbc, 0x7c, 0x06, 0xff, 0x1e, 0x5d, 0xfd, 0x92, 0x06, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x24, 0x4c, 0x29, 0x57, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	PostCreate(ctx context.Context, in *Post, opts ...grpc.CallOption) (*OkBOOL, error)
	PostGetByID(ctx context.Context, in *GetIdFromUser, opts ...grpc.CallOption) (*Post, error)
	PostDeleteByID(ctx context.Context, in *GetIdFromUser, opts ...grpc.CallOption) (*OkBOOL, error)
	PostGetAllPosts(ctx context.Context, in *GetIdFromUser, opts ...grpc.CallOption) (*AllPost, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) PostCreate(ctx context.Context, in *Post, opts ...grpc.CallOption) (*OkBOOL, error) {
	out := new(OkBOOL)
	err := c.cc.Invoke(ctx, "/user.PostService/PostCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostGetByID(ctx context.Context, in *GetIdFromUser, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/user.PostService/PostGetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostDeleteByID(ctx context.Context, in *GetIdFromUser, opts ...grpc.CallOption) (*OkBOOL, error) {
	out := new(OkBOOL)
	err := c.cc.Invoke(ctx, "/user.PostService/PostDeleteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostGetAllPosts(ctx context.Context, in *GetIdFromUser, opts ...grpc.CallOption) (*AllPost, error) {
	out := new(AllPost)
	err := c.cc.Invoke(ctx, "/user.PostService/PostGetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	PostCreate(context.Context, *Post) (*OkBOOL, error)
	PostGetByID(context.Context, *GetIdFromUser) (*Post, error)
	PostDeleteByID(context.Context, *GetIdFromUser) (*OkBOOL, error)
	PostGetAllPosts(context.Context, *GetIdFromUser) (*AllPost, error)
}

// UnimplementedPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (*UnimplementedPostServiceServer) PostCreate(ctx context.Context, req *Post) (*OkBOOL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCreate not implemented")
}
func (*UnimplementedPostServiceServer) PostGetByID(ctx context.Context, req *GetIdFromUser) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGetByID not implemented")
}
func (*UnimplementedPostServiceServer) PostDeleteByID(ctx context.Context, req *GetIdFromUser) (*OkBOOL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDeleteByID not implemented")
}
func (*UnimplementedPostServiceServer) PostGetAllPosts(ctx context.Context, req *GetIdFromUser) (*AllPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGetAllPosts not implemented")
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_PostCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.PostService/PostCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostCreate(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostGetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdFromUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostGetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.PostService/PostGetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostGetByID(ctx, req.(*GetIdFromUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostDeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdFromUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostDeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.PostService/PostDeleteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostDeleteByID(ctx, req.(*GetIdFromUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostGetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdFromUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostGetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.PostService/PostGetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostGetAllPosts(ctx, req.(*GetIdFromUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostCreate",
			Handler:    _PostService_PostCreate_Handler,
		},
		{
			MethodName: "PostGetByID",
			Handler:    _PostService_PostGetByID_Handler,
		},
		{
			MethodName: "PostDeleteByID",
			Handler:    _PostService_PostDeleteByID_Handler,
		},
		{
			MethodName: "PostGetAllPosts",
			Handler:    _PostService_PostGetAllPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}

func (m *AllPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllPost) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Posts) > 0 {
		for _, msg := range m.Posts {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPost(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Post) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Post) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.Medias) > 0 {
		for _, msg := range m.Medias {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintPost(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Media) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Media) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.PostId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.PostId)))
		i += copy(dAtA[i:], m.PostId)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Link) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Link)))
		i += copy(dAtA[i:], m.Link)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetIdFromUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetIdFromUser) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPost(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *OkBOOL) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OkBOOL) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status {
		dAtA[i] = 0x8
		i++
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPost(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AllPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovPost(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	if len(m.Medias) > 0 {
		for _, e := range m.Medias {
			l = e.Size()
			n += 1 + l + sovPost(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Media) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Link)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetIdFromUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OkBOOL) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPost(x uint64) (n int) {
	return sovPost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Posts = append(m.Posts, &Post{})
			if err := m.Posts[len(m.Posts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Post) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Post: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Post: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Medias", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Medias = append(m.Medias, &Media{})
			if err := m.Medias[len(m.Medias)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Media) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Media: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Media: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Link = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetIdFromUser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetIdFromUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetIdFromUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OkBOOL) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OkBOOL: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OkBOOL: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPost
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPost
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPost
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPost
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPost(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPost
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPost = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPost   = fmt.Errorf("proto: integer overflow")
)
