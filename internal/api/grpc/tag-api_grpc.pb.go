// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TagClient is the client API for Tag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagClient interface {
	GetAuthorizedCards(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAuthorizedCardsResponse, error)
	AddAuthorizedCards(ctx context.Context, in *AddAuthorizedCardsRequest, opts ...grpc.CallOption) (*AddAuthorizedCardsResponse, error)
	RemoveAuthorizedCard(ctx context.Context, in *AddAuthorizedCardsRequest, opts ...grpc.CallOption) (*AddAuthorizedCardsResponse, error)
}

type tagClient struct {
	cc grpc.ClientConnInterface
}

func NewTagClient(cc grpc.ClientConnInterface) TagClient {
	return &tagClient{cc}
}

func (c *tagClient) GetAuthorizedCards(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAuthorizedCardsResponse, error) {
	out := new(GetAuthorizedCardsResponse)
	err := c.cc.Invoke(ctx, "/api.Tag/GetAuthorizedCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) AddAuthorizedCards(ctx context.Context, in *AddAuthorizedCardsRequest, opts ...grpc.CallOption) (*AddAuthorizedCardsResponse, error) {
	out := new(AddAuthorizedCardsResponse)
	err := c.cc.Invoke(ctx, "/api.Tag/AddAuthorizedCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) RemoveAuthorizedCard(ctx context.Context, in *AddAuthorizedCardsRequest, opts ...grpc.CallOption) (*AddAuthorizedCardsResponse, error) {
	out := new(AddAuthorizedCardsResponse)
	err := c.cc.Invoke(ctx, "/api.Tag/RemoveAuthorizedCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServer is the server API for Tag service.
// All implementations must embed UnimplementedTagServer
// for forward compatibility
type TagServer interface {
	GetAuthorizedCards(context.Context, *empty.Empty) (*GetAuthorizedCardsResponse, error)
	AddAuthorizedCards(context.Context, *AddAuthorizedCardsRequest) (*AddAuthorizedCardsResponse, error)
	RemoveAuthorizedCard(context.Context, *AddAuthorizedCardsRequest) (*AddAuthorizedCardsResponse, error)
	mustEmbedUnimplementedTagServer()
}

// UnimplementedTagServer must be embedded to have forward compatible implementations.
type UnimplementedTagServer struct {
}

func (UnimplementedTagServer) GetAuthorizedCards(context.Context, *empty.Empty) (*GetAuthorizedCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizedCards not implemented")
}
func (UnimplementedTagServer) AddAuthorizedCards(context.Context, *AddAuthorizedCardsRequest) (*AddAuthorizedCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthorizedCards not implemented")
}
func (UnimplementedTagServer) RemoveAuthorizedCard(context.Context, *AddAuthorizedCardsRequest) (*AddAuthorizedCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAuthorizedCard not implemented")
}
func (UnimplementedTagServer) mustEmbedUnimplementedTagServer() {}

// UnsafeTagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServer will
// result in compilation errors.
type UnsafeTagServer interface {
	mustEmbedUnimplementedTagServer()
}

func RegisterTagServer(s grpc.ServiceRegistrar, srv TagServer) {
	s.RegisterService(&Tag_ServiceDesc, srv)
}

func _Tag_GetAuthorizedCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).GetAuthorizedCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Tag/GetAuthorizedCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).GetAuthorizedCards(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_AddAuthorizedCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAuthorizedCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).AddAuthorizedCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Tag/AddAuthorizedCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).AddAuthorizedCards(ctx, req.(*AddAuthorizedCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_RemoveAuthorizedCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAuthorizedCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).RemoveAuthorizedCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Tag/RemoveAuthorizedCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).RemoveAuthorizedCard(ctx, req.(*AddAuthorizedCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tag_ServiceDesc is the grpc.ServiceDesc for Tag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Tag",
	HandlerType: (*TagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthorizedCards",
			Handler:    _Tag_GetAuthorizedCards_Handler,
		},
		{
			MethodName: "AddAuthorizedCards",
			Handler:    _Tag_AddAuthorizedCards_Handler,
		},
		{
			MethodName: "RemoveAuthorizedCard",
			Handler:    _Tag_RemoveAuthorizedCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tag-api.proto",
}