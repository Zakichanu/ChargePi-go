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

// EvseClient is the client API for Evse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvseClient interface {
	AddEVSE(ctx context.Context, in *SetEVCCRequest, opts ...grpc.CallOption) (*AddEvseResponse, error)
	GetEVSEs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEvsesResponse, error)
	GetEVSE(ctx context.Context, in *GetEvseRequest, opts ...grpc.CallOption) (*GetEvseResponse, error)
	SetEVCC(ctx context.Context, in *SetEVCCRequest, opts ...grpc.CallOption) (*SetEvccResponse, error)
	SetPowerMeter(ctx context.Context, in *SetPowerMeterRequest, opts ...grpc.CallOption) (*SetPowerMeterDetails, error)
	GetUsageForEVSE(ctx context.Context, in *GetUsageForEVSERequest, opts ...grpc.CallOption) (Evse_GetUsageForEVSEClient, error)
}

type evseClient struct {
	cc grpc.ClientConnInterface
}

func NewEvseClient(cc grpc.ClientConnInterface) EvseClient {
	return &evseClient{cc}
}

func (c *evseClient) AddEVSE(ctx context.Context, in *SetEVCCRequest, opts ...grpc.CallOption) (*AddEvseResponse, error) {
	out := new(AddEvseResponse)
	err := c.cc.Invoke(ctx, "/api.Evse/AddEVSE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evseClient) GetEVSEs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEvsesResponse, error) {
	out := new(GetEvsesResponse)
	err := c.cc.Invoke(ctx, "/api.Evse/GetEVSEs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evseClient) GetEVSE(ctx context.Context, in *GetEvseRequest, opts ...grpc.CallOption) (*GetEvseResponse, error) {
	out := new(GetEvseResponse)
	err := c.cc.Invoke(ctx, "/api.Evse/GetEVSE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evseClient) SetEVCC(ctx context.Context, in *SetEVCCRequest, opts ...grpc.CallOption) (*SetEvccResponse, error) {
	out := new(SetEvccResponse)
	err := c.cc.Invoke(ctx, "/api.Evse/SetEVCC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evseClient) SetPowerMeter(ctx context.Context, in *SetPowerMeterRequest, opts ...grpc.CallOption) (*SetPowerMeterDetails, error) {
	out := new(SetPowerMeterDetails)
	err := c.cc.Invoke(ctx, "/api.Evse/SetPowerMeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evseClient) GetUsageForEVSE(ctx context.Context, in *GetUsageForEVSERequest, opts ...grpc.CallOption) (Evse_GetUsageForEVSEClient, error) {
	stream, err := c.cc.NewStream(ctx, &Evse_ServiceDesc.Streams[0], "/api.Evse/GetUsageForEVSE", opts...)
	if err != nil {
		return nil, err
	}
	x := &evseGetUsageForEVSEClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Evse_GetUsageForEVSEClient interface {
	Recv() (*GetUsageForEVSEResponse, error)
	grpc.ClientStream
}

type evseGetUsageForEVSEClient struct {
	grpc.ClientStream
}

func (x *evseGetUsageForEVSEClient) Recv() (*GetUsageForEVSEResponse, error) {
	m := new(GetUsageForEVSEResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EvseServer is the server API for Evse service.
// All implementations must embed UnimplementedEvseServer
// for forward compatibility
type EvseServer interface {
	AddEVSE(context.Context, *SetEVCCRequest) (*AddEvseResponse, error)
	GetEVSEs(context.Context, *empty.Empty) (*GetEvsesResponse, error)
	GetEVSE(context.Context, *GetEvseRequest) (*GetEvseResponse, error)
	SetEVCC(context.Context, *SetEVCCRequest) (*SetEvccResponse, error)
	SetPowerMeter(context.Context, *SetPowerMeterRequest) (*SetPowerMeterDetails, error)
	GetUsageForEVSE(*GetUsageForEVSERequest, Evse_GetUsageForEVSEServer) error
	mustEmbedUnimplementedEvseServer()
}

// UnimplementedEvseServer must be embedded to have forward compatible implementations.
type UnimplementedEvseServer struct {
}

func (UnimplementedEvseServer) AddEVSE(context.Context, *SetEVCCRequest) (*AddEvseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEVSE not implemented")
}
func (UnimplementedEvseServer) GetEVSEs(context.Context, *empty.Empty) (*GetEvsesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEVSEs not implemented")
}
func (UnimplementedEvseServer) GetEVSE(context.Context, *GetEvseRequest) (*GetEvseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEVSE not implemented")
}
func (UnimplementedEvseServer) SetEVCC(context.Context, *SetEVCCRequest) (*SetEvccResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEVCC not implemented")
}
func (UnimplementedEvseServer) SetPowerMeter(context.Context, *SetPowerMeterRequest) (*SetPowerMeterDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPowerMeter not implemented")
}
func (UnimplementedEvseServer) GetUsageForEVSE(*GetUsageForEVSERequest, Evse_GetUsageForEVSEServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsageForEVSE not implemented")
}
func (UnimplementedEvseServer) mustEmbedUnimplementedEvseServer() {}

// UnsafeEvseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvseServer will
// result in compilation errors.
type UnsafeEvseServer interface {
	mustEmbedUnimplementedEvseServer()
}

func RegisterEvseServer(s grpc.ServiceRegistrar, srv EvseServer) {
	s.RegisterService(&Evse_ServiceDesc, srv)
}

func _Evse_AddEVSE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEVCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvseServer).AddEVSE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evse/AddEVSE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvseServer).AddEVSE(ctx, req.(*SetEVCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evse_GetEVSEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvseServer).GetEVSEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evse/GetEVSEs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvseServer).GetEVSEs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evse_GetEVSE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEvseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvseServer).GetEVSE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evse/GetEVSE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvseServer).GetEVSE(ctx, req.(*GetEvseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evse_SetEVCC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEVCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvseServer).SetEVCC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evse/SetEVCC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvseServer).SetEVCC(ctx, req.(*SetEVCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evse_SetPowerMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPowerMeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvseServer).SetPowerMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evse/SetPowerMeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvseServer).SetPowerMeter(ctx, req.(*SetPowerMeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Evse_GetUsageForEVSE_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUsageForEVSERequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EvseServer).GetUsageForEVSE(m, &evseGetUsageForEVSEServer{stream})
}

type Evse_GetUsageForEVSEServer interface {
	Send(*GetUsageForEVSEResponse) error
	grpc.ServerStream
}

type evseGetUsageForEVSEServer struct {
	grpc.ServerStream
}

func (x *evseGetUsageForEVSEServer) Send(m *GetUsageForEVSEResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Evse_ServiceDesc is the grpc.ServiceDesc for Evse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Evse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Evse",
	HandlerType: (*EvseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEVSE",
			Handler:    _Evse_AddEVSE_Handler,
		},
		{
			MethodName: "GetEVSEs",
			Handler:    _Evse_GetEVSEs_Handler,
		},
		{
			MethodName: "GetEVSE",
			Handler:    _Evse_GetEVSE_Handler,
		},
		{
			MethodName: "SetEVCC",
			Handler:    _Evse_SetEVCC_Handler,
		},
		{
			MethodName: "SetPowerMeter",
			Handler:    _Evse_SetPowerMeter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsageForEVSE",
			Handler:       _Evse_GetUsageForEVSE_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "evse-api.proto",
}
