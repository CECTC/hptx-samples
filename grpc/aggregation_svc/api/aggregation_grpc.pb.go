// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.4
// source: aggregation.proto

package api

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

// AggregationServiceClient is the client API for AggregationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregationServiceClient interface {
	CreateSoCommit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error)
	CreateSoRollback(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error)
}

type aggregationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregationServiceClient(cc grpc.ClientConnInterface) AggregationServiceClient {
	return &aggregationServiceClient{cc}
}

func (c *aggregationServiceClient) CreateSoCommit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.AggregationService/CreateSoCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregationServiceClient) CreateSoRollback(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.AggregationService/CreateSoRollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregationServiceServer is the server API for AggregationService service.
// All implementations should embed UnimplementedAggregationServiceServer
// for forward compatibility
type AggregationServiceServer interface {
	CreateSoCommit(context.Context, *empty.Empty) (*Response, error)
	CreateSoRollback(context.Context, *empty.Empty) (*Response, error)
}

// UnimplementedAggregationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAggregationServiceServer struct {
}

func (UnimplementedAggregationServiceServer) CreateSoCommit(context.Context, *empty.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSoCommit not implemented")
}
func (UnimplementedAggregationServiceServer) CreateSoRollback(context.Context, *empty.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSoRollback not implemented")
}

// UnsafeAggregationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregationServiceServer will
// result in compilation errors.
type UnsafeAggregationServiceServer interface {
	mustEmbedUnimplementedAggregationServiceServer()
}

func RegisterAggregationServiceServer(s grpc.ServiceRegistrar, srv AggregationServiceServer) {
	s.RegisterService(&AggregationService_ServiceDesc, srv)
}

func _AggregationService_CreateSoCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregationServiceServer).CreateSoCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AggregationService/CreateSoCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregationServiceServer).CreateSoCommit(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregationService_CreateSoRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregationServiceServer).CreateSoRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AggregationService/CreateSoRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregationServiceServer).CreateSoRollback(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregationService_ServiceDesc is the grpc.ServiceDesc for AggregationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AggregationService",
	HandlerType: (*AggregationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSoCommit",
			Handler:    _AggregationService_CreateSoCommit_Handler,
		},
		{
			MethodName: "CreateSoRollback",
			Handler:    _AggregationService_CreateSoRollback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aggregation.proto",
}