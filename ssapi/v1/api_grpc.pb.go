// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: ssapi/v1/api.proto

package ssapiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ServerSideApi_Publish_FullMethodName     = "/edgehub.protocol.ssapi.v1.ServerSideApi/Publish"
	ServerSideApi_Survey_FullMethodName      = "/edgehub.protocol.ssapi.v1.ServerSideApi/Survey"
	ServerSideApi_Subscribe_FullMethodName   = "/edgehub.protocol.ssapi.v1.ServerSideApi/Subscribe"
	ServerSideApi_Unsubscribe_FullMethodName = "/edgehub.protocol.ssapi.v1.ServerSideApi/Unsubscribe"
	ServerSideApi_Disconnect_FullMethodName  = "/edgehub.protocol.ssapi.v1.ServerSideApi/Disconnect"
)

// ServerSideApiClient is the client API for ServerSideApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerSideApiClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error)
	Survey(ctx context.Context, in *SurveyRequest, opts ...grpc.CallOption) (*SurveyReply, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeReply, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectReply, error)
}

type serverSideApiClient struct {
	cc grpc.ClientConnInterface
}

func NewServerSideApiClient(cc grpc.ClientConnInterface) ServerSideApiClient {
	return &serverSideApiClient{cc}
}

func (c *serverSideApiClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error) {
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, ServerSideApi_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverSideApiClient) Survey(ctx context.Context, in *SurveyRequest, opts ...grpc.CallOption) (*SurveyReply, error) {
	out := new(SurveyReply)
	err := c.cc.Invoke(ctx, ServerSideApi_Survey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverSideApiClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error) {
	out := new(SubscribeReply)
	err := c.cc.Invoke(ctx, ServerSideApi_Subscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverSideApiClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeReply, error) {
	out := new(UnsubscribeReply)
	err := c.cc.Invoke(ctx, ServerSideApi_Unsubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverSideApiClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectReply, error) {
	out := new(DisconnectReply)
	err := c.cc.Invoke(ctx, ServerSideApi_Disconnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerSideApiServer is the server API for ServerSideApi service.
// All implementations must embed UnimplementedServerSideApiServer
// for forward compatibility
type ServerSideApiServer interface {
	Publish(context.Context, *PublishRequest) (*PublishReply, error)
	Survey(context.Context, *SurveyRequest) (*SurveyReply, error)
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeReply, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeReply, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectReply, error)
	mustEmbedUnimplementedServerSideApiServer()
}

// UnimplementedServerSideApiServer must be embedded to have forward compatible implementations.
type UnimplementedServerSideApiServer struct {
}

func (UnimplementedServerSideApiServer) Publish(context.Context, *PublishRequest) (*PublishReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedServerSideApiServer) Survey(context.Context, *SurveyRequest) (*SurveyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Survey not implemented")
}
func (UnimplementedServerSideApiServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedServerSideApiServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedServerSideApiServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedServerSideApiServer) mustEmbedUnimplementedServerSideApiServer() {}

// UnsafeServerSideApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerSideApiServer will
// result in compilation errors.
type UnsafeServerSideApiServer interface {
	mustEmbedUnimplementedServerSideApiServer()
}

func RegisterServerSideApiServer(s grpc.ServiceRegistrar, srv ServerSideApiServer) {
	s.RegisterService(&ServerSideApi_ServiceDesc, srv)
}

func _ServerSideApi_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSideApiServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerSideApi_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSideApiServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerSideApi_Survey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SurveyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSideApiServer).Survey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerSideApi_Survey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSideApiServer).Survey(ctx, req.(*SurveyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerSideApi_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSideApiServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerSideApi_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSideApiServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerSideApi_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSideApiServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerSideApi_Unsubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSideApiServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerSideApi_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSideApiServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerSideApi_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSideApiServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerSideApi_ServiceDesc is the grpc.ServiceDesc for ServerSideApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerSideApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edgehub.protocol.ssapi.v1.ServerSideApi",
	HandlerType: (*ServerSideApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _ServerSideApi_Publish_Handler,
		},
		{
			MethodName: "Survey",
			Handler:    _ServerSideApi_Survey_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _ServerSideApi_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _ServerSideApi_Unsubscribe_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _ServerSideApi_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ssapi/v1/api.proto",
}
