// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/schema.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DistributedCfgServiceMKClient is the client API for DistributedCfgServiceMK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistributedCfgServiceMKClient interface {
	// config basic CRUD calls
	CreateConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Timestamp, error)
	UpdateConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Timestamp, error)
	GetConfig(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Config, error)
	DeleteConfig(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Timestamp, error)
	// config version history related calls:
	GetArchivedConfig(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*ConfigByTimestamp, error)
	ListConfigTimestamps(ctx context.Context, in *Service, opts ...grpc.CallOption) (*TimestampList, error)
	// config blocking for deletion by subscribed apps related calls:
	SubscribeClientApp(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnSubscribeClientApp(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListConfigSubscribers(ctx context.Context, in *Service, opts ...grpc.CallOption) (*ConfigSubscribers, error)
}

type distributedCfgServiceMKClient struct {
	cc grpc.ClientConnInterface
}

func NewDistributedCfgServiceMKClient(cc grpc.ClientConnInterface) DistributedCfgServiceMKClient {
	return &distributedCfgServiceMKClient{cc}
}

func (c *distributedCfgServiceMKClient) CreateConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Timestamp, error) {
	out := new(Timestamp)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/CreateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) UpdateConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Timestamp, error) {
	out := new(Timestamp)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) GetConfig(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) DeleteConfig(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Timestamp, error) {
	out := new(Timestamp)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) GetArchivedConfig(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*ConfigByTimestamp, error) {
	out := new(ConfigByTimestamp)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/GetArchivedConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) ListConfigTimestamps(ctx context.Context, in *Service, opts ...grpc.CallOption) (*TimestampList, error) {
	out := new(TimestampList)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/ListConfigTimestamps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) SubscribeClientApp(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/SubscribeClientApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) UnSubscribeClientApp(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/UnSubscribeClientApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedCfgServiceMKClient) ListConfigSubscribers(ctx context.Context, in *Service, opts ...grpc.CallOption) (*ConfigSubscribers, error) {
	out := new(ConfigSubscribers)
	err := c.cc.Invoke(ctx, "/distributed_cfg_service_mk.DistributedCfgServiceMK/ListConfigSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributedCfgServiceMKServer is the server API for DistributedCfgServiceMK service.
// All implementations must embed UnimplementedDistributedCfgServiceMKServer
// for forward compatibility
type DistributedCfgServiceMKServer interface {
	// config basic CRUD calls
	CreateConfig(context.Context, *Config) (*Timestamp, error)
	UpdateConfig(context.Context, *Config) (*Timestamp, error)
	GetConfig(context.Context, *Service) (*Config, error)
	DeleteConfig(context.Context, *Service) (*Timestamp, error)
	// config version history related calls:
	GetArchivedConfig(context.Context, *Timestamp) (*ConfigByTimestamp, error)
	ListConfigTimestamps(context.Context, *Service) (*TimestampList, error)
	// config blocking for deletion by subscribed apps related calls:
	SubscribeClientApp(context.Context, *SubscriptionRequest) (*emptypb.Empty, error)
	UnSubscribeClientApp(context.Context, *SubscriptionRequest) (*emptypb.Empty, error)
	ListConfigSubscribers(context.Context, *Service) (*ConfigSubscribers, error)
	mustEmbedUnimplementedDistributedCfgServiceMKServer()
}

// UnimplementedDistributedCfgServiceMKServer must be embedded to have forward compatible implementations.
type UnimplementedDistributedCfgServiceMKServer struct {
}

func (UnimplementedDistributedCfgServiceMKServer) CreateConfig(context.Context, *Config) (*Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfig not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) UpdateConfig(context.Context, *Config) (*Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) GetConfig(context.Context, *Service) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) DeleteConfig(context.Context, *Service) (*Timestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) GetArchivedConfig(context.Context, *Timestamp) (*ConfigByTimestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArchivedConfig not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) ListConfigTimestamps(context.Context, *Service) (*TimestampList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigTimestamps not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) SubscribeClientApp(context.Context, *SubscriptionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeClientApp not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) UnSubscribeClientApp(context.Context, *SubscriptionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribeClientApp not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) ListConfigSubscribers(context.Context, *Service) (*ConfigSubscribers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigSubscribers not implemented")
}
func (UnimplementedDistributedCfgServiceMKServer) mustEmbedUnimplementedDistributedCfgServiceMKServer() {
}

// UnsafeDistributedCfgServiceMKServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistributedCfgServiceMKServer will
// result in compilation errors.
type UnsafeDistributedCfgServiceMKServer interface {
	mustEmbedUnimplementedDistributedCfgServiceMKServer()
}

func RegisterDistributedCfgServiceMKServer(s grpc.ServiceRegistrar, srv DistributedCfgServiceMKServer) {
	s.RegisterService(&DistributedCfgServiceMK_ServiceDesc, srv)
}

func _DistributedCfgServiceMK_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/CreateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).CreateConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).UpdateConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).GetConfig(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).DeleteConfig(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_GetArchivedConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timestamp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).GetArchivedConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/GetArchivedConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).GetArchivedConfig(ctx, req.(*Timestamp))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_ListConfigTimestamps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).ListConfigTimestamps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/ListConfigTimestamps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).ListConfigTimestamps(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_SubscribeClientApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).SubscribeClientApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/SubscribeClientApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).SubscribeClientApp(ctx, req.(*SubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_UnSubscribeClientApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).UnSubscribeClientApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/UnSubscribeClientApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).UnSubscribeClientApp(ctx, req.(*SubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedCfgServiceMK_ListConfigSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedCfgServiceMKServer).ListConfigSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributed_cfg_service_mk.DistributedCfgServiceMK/ListConfigSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedCfgServiceMKServer).ListConfigSubscribers(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

// DistributedCfgServiceMK_ServiceDesc is the grpc.ServiceDesc for DistributedCfgServiceMK service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DistributedCfgServiceMK_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "distributed_cfg_service_mk.DistributedCfgServiceMK",
	HandlerType: (*DistributedCfgServiceMKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConfig",
			Handler:    _DistributedCfgServiceMK_CreateConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _DistributedCfgServiceMK_UpdateConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _DistributedCfgServiceMK_GetConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _DistributedCfgServiceMK_DeleteConfig_Handler,
		},
		{
			MethodName: "GetArchivedConfig",
			Handler:    _DistributedCfgServiceMK_GetArchivedConfig_Handler,
		},
		{
			MethodName: "ListConfigTimestamps",
			Handler:    _DistributedCfgServiceMK_ListConfigTimestamps_Handler,
		},
		{
			MethodName: "SubscribeClientApp",
			Handler:    _DistributedCfgServiceMK_SubscribeClientApp_Handler,
		},
		{
			MethodName: "UnSubscribeClientApp",
			Handler:    _DistributedCfgServiceMK_UnSubscribeClientApp_Handler,
		},
		{
			MethodName: "ListConfigSubscribers",
			Handler:    _DistributedCfgServiceMK_ListConfigSubscribers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/schema.proto",
}
