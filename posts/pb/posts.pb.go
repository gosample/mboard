// Code generated by protoc-gen-go.
// source: pb/posts.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/posts.proto

It has these top-level messages:
	Post
	GetLatestPostsRequest
	AddPostRequest
	GetPostRequest
	GetPostResponse
	GetLatestPostsResponse
	AddPostResponse
*/
package pb

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

type Post struct {
	Id     []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Text   string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	Date   int64  `protobuf:"varint,4,opt,name=date" json:"date,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Post) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Post) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Post) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Post) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type GetLatestPostsRequest struct {
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetLatestPostsRequest) Reset()                    { *m = GetLatestPostsRequest{} }
func (m *GetLatestPostsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLatestPostsRequest) ProtoMessage()               {}
func (*GetLatestPostsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetLatestPostsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetLatestPostsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type AddPostRequest struct {
	Author string `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *AddPostRequest) Reset()                    { *m = AddPostRequest{} }
func (m *AddPostRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPostRequest) ProtoMessage()               {}
func (*AddPostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddPostRequest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *AddPostRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type GetPostRequest struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetPostRequest) Reset()                    { *m = GetPostRequest{} }
func (m *GetPostRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPostRequest) ProtoMessage()               {}
func (*GetPostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetPostRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetPostResponse struct {
	Post *Post  `protobuf:"bytes,1,opt,name=post" json:"post,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetPostResponse) Reset()                    { *m = GetPostResponse{} }
func (m *GetPostResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPostResponse) ProtoMessage()               {}
func (*GetPostResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetPostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *GetPostResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetLatestPostsResponse struct {
	Posts []*Post `protobuf:"bytes,1,rep,name=posts" json:"posts,omitempty"`
	Err   string  `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetLatestPostsResponse) Reset()                    { *m = GetLatestPostsResponse{} }
func (m *GetLatestPostsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLatestPostsResponse) ProtoMessage()               {}
func (*GetLatestPostsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetLatestPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GetLatestPostsResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type AddPostResponse struct {
	Id  []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *AddPostResponse) Reset()                    { *m = AddPostResponse{} }
func (m *AddPostResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPostResponse) ProtoMessage()               {}
func (*AddPostResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddPostResponse) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AddPostResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "pb.Post")
	proto.RegisterType((*GetLatestPostsRequest)(nil), "pb.GetLatestPostsRequest")
	proto.RegisterType((*AddPostRequest)(nil), "pb.AddPostRequest")
	proto.RegisterType((*GetPostRequest)(nil), "pb.GetPostRequest")
	proto.RegisterType((*GetPostResponse)(nil), "pb.GetPostResponse")
	proto.RegisterType((*GetLatestPostsResponse)(nil), "pb.GetLatestPostsResponse")
	proto.RegisterType((*AddPostResponse)(nil), "pb.AddPostResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Posts service

type PostsClient interface {
	GetLatestPosts(ctx context.Context, in *GetLatestPostsRequest, opts ...grpc.CallOption) (*GetLatestPostsResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*AddPostResponse, error)
}

type postsClient struct {
	cc *grpc.ClientConn
}

func NewPostsClient(cc *grpc.ClientConn) PostsClient {
	return &postsClient{cc}
}

func (c *postsClient) GetLatestPosts(ctx context.Context, in *GetLatestPostsRequest, opts ...grpc.CallOption) (*GetLatestPostsResponse, error) {
	out := new(GetLatestPostsResponse)
	err := grpc.Invoke(ctx, "/pb.Posts/GetLatestPosts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := grpc.Invoke(ctx, "/pb.Posts/GetPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*AddPostResponse, error) {
	out := new(AddPostResponse)
	err := grpc.Invoke(ctx, "/pb.Posts/AddPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Posts service

type PostsServer interface {
	GetLatestPosts(context.Context, *GetLatestPostsRequest) (*GetLatestPostsResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	AddPost(context.Context, *AddPostRequest) (*AddPostResponse, error)
}

func RegisterPostsServer(s *grpc.Server, srv PostsServer) {
	s.RegisterService(&_Posts_serviceDesc, srv)
}

func _Posts_GetLatestPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetLatestPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Posts/GetLatestPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetLatestPosts(ctx, req.(*GetLatestPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Posts/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_AddPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).AddPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Posts/AddPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).AddPost(ctx, req.(*AddPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Posts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Posts",
	HandlerType: (*PostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestPosts",
			Handler:    _Posts_GetLatestPosts_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Posts_GetPost_Handler,
		},
		{
			MethodName: "AddPost",
			Handler:    _Posts_AddPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/posts.proto",
}

func init() { proto.RegisterFile("pb/posts.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x4e, 0xc3, 0x30,
	0x0c, 0x5d, 0xda, 0x6e, 0x80, 0x41, 0x1d, 0x32, 0x30, 0x95, 0x09, 0xa1, 0x2a, 0xa7, 0x9d, 0x86,
	0xb4, 0x71, 0xe4, 0xb2, 0x03, 0x9a, 0x40, 0x1c, 0x50, 0x0e, 0xdc, 0x57, 0x35, 0x13, 0x95, 0x80,
	0x84, 0xc5, 0x93, 0xf8, 0x3a, 0xbe, 0x0d, 0xc5, 0x0d, 0xa3, 0x2b, 0xbd, 0xd9, 0x8e, 0xdf, 0x7b,
	0xf6, 0x73, 0x20, 0xb5, 0xc5, 0x8d, 0x35, 0x8e, 0xdc, 0xd4, 0x6e, 0x0c, 0x19, 0x8c, 0x6c, 0x21,
	0x5f, 0x20, 0x79, 0x36, 0x8e, 0x30, 0x85, 0xa8, 0x2a, 0x33, 0x91, 0x8b, 0xc9, 0x89, 0x8a, 0xaa,
	0x12, 0x47, 0x30, 0x58, 0x6d, 0xe9, 0xd5, 0x6c, 0xb2, 0x28, 0x17, 0x93, 0x23, 0x15, 0x32, 0x44,
	0x48, 0x48, 0x7f, 0x51, 0x16, 0x73, 0x95, 0x63, 0x5f, 0x2b, 0x57, 0xa4, 0xb3, 0x24, 0x17, 0x93,
	0x58, 0x71, 0x2c, 0xef, 0xe1, 0x62, 0xa9, 0xe9, 0x69, 0x45, 0xda, 0x91, 0x17, 0x70, 0x4a, 0x7f,
	0x6e, 0xb5, 0x23, 0x4f, 0x6c, 0xd6, 0x6b, 0xa7, 0x89, 0xc5, 0x62, 0x15, 0x32, 0x3c, 0x87, 0xfe,
	0x5b, 0xf5, 0x5e, 0x11, 0xeb, 0xc5, 0xaa, 0x4e, 0xe4, 0x1d, 0xa4, 0x8b, 0xb2, 0xf4, 0x04, 0x0d,
	0x7c, 0x18, 0x4c, 0x74, 0x0e, 0x16, 0xfd, 0x0d, 0x26, 0x73, 0x48, 0x97, 0x9a, 0x9a, 0xe8, 0xd6,
	0x9a, 0x72, 0x01, 0xc3, 0x5d, 0x87, 0xb3, 0xe6, 0xc3, 0x69, 0xbc, 0x82, 0xc4, 0x9b, 0xc4, 0x4d,
	0xc7, 0xb3, 0xc3, 0xa9, 0x2d, 0xa6, 0xfc, 0xce, 0x55, 0x3c, 0x85, 0x58, 0x6f, 0x7e, 0x4d, 0xf1,
	0xa1, 0x7c, 0x84, 0x51, 0x7b, 0xd3, 0xc0, 0x74, 0x0d, 0x7d, 0xb6, 0x3b, 0x13, 0x79, 0xbc, 0x47,
	0x55, 0x97, 0x3b, 0xb8, 0xe6, 0x30, 0xdc, 0xad, 0x1b, 0x48, 0xda, 0x87, 0xf9, 0x07, 0x9a, 0x7d,
	0x0b, 0xe8, 0xb3, 0x30, 0x3e, 0xf0, 0xbe, 0x8d, 0x51, 0xf0, 0xd2, 0x6b, 0x76, 0x1e, 0x62, 0x3c,
	0xee, 0x7a, 0xaa, 0x45, 0x65, 0x0f, 0x6f, 0xe1, 0x20, 0x18, 0x83, 0x18, 0x1a, 0x1b, 0x3e, 0x8e,
	0xcf, 0xf6, 0x6a, 0x4d, 0x54, 0x98, 0xbf, 0x46, 0xed, 0xdf, 0xae, 0x46, 0xb5, 0x16, 0x94, 0xbd,
	0x62, 0xc0, 0xdf, 0x71, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xda, 0xba, 0x26, 0xa0, 0x02,
	0x00, 0x00,
}
