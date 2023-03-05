// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: monolith.proto

package monolith

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

// MonolithClient is the client API for Monolith service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonolithClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	IngestFiles(ctx context.Context, in *IngestFilesRequest, opts ...grpc.CallOption) (*IngestFilesResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type monolithClient struct {
	cc grpc.ClientConnInterface
}

func NewMonolithClient(cc grpc.ClientConnInterface) MonolithClient {
	return &monolithClient{cc}
}

func (c *monolithClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/monolith.Monolith/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monolithClient) IngestFiles(ctx context.Context, in *IngestFilesRequest, opts ...grpc.CallOption) (*IngestFilesResponse, error) {
	out := new(IngestFilesResponse)
	err := c.cc.Invoke(ctx, "/monolith.Monolith/IngestFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monolithClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := c.cc.Invoke(ctx, "/monolith.Monolith/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monolithClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/monolith.Monolith/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonolithServer is the server API for Monolith service.
// All implementations must embed UnimplementedMonolithServer
// for forward compatibility
type MonolithServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	IngestFiles(context.Context, *IngestFilesRequest) (*IngestFilesResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedMonolithServer()
}

// UnimplementedMonolithServer must be embedded to have forward compatible implementations.
type UnimplementedMonolithServer struct {
}

func (UnimplementedMonolithServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedMonolithServer) IngestFiles(context.Context, *IngestFilesRequest) (*IngestFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngestFiles not implemented")
}
func (UnimplementedMonolithServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedMonolithServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedMonolithServer) mustEmbedUnimplementedMonolithServer() {}

// UnsafeMonolithServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonolithServer will
// result in compilation errors.
type UnsafeMonolithServer interface {
	mustEmbedUnimplementedMonolithServer()
}

func RegisterMonolithServer(s grpc.ServiceRegistrar, srv MonolithServer) {
	s.RegisterService(&Monolith_ServiceDesc, srv)
}

func _Monolith_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonolithServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monolith.Monolith/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonolithServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monolith_IngestFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonolithServer).IngestFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monolith.Monolith/IngestFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonolithServer).IngestFiles(ctx, req.(*IngestFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monolith_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonolithServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monolith.Monolith/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonolithServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monolith_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonolithServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monolith.Monolith/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonolithServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Monolith_ServiceDesc is the grpc.ServiceDesc for Monolith service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monolith_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monolith.Monolith",
	HandlerType: (*MonolithServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Monolith_Hello_Handler,
		},
		{
			MethodName: "IngestFiles",
			Handler:    _Monolith_IngestFiles_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _Monolith_DeleteIndex_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Monolith_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monolith.proto",
}