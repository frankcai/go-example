// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: documentsvc.proto

package documentsvc

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

// DocumentSvcClient is the client API for DocumentSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentSvcClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type documentSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentSvcClient(cc grpc.ClientConnInterface) DocumentSvcClient {
	return &documentSvcClient{cc}
}

func (c *documentSvcClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/documentsvc.DocumentSvc/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentSvcClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := c.cc.Invoke(ctx, "/documentsvc.DocumentSvc/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentSvcClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/documentsvc.DocumentSvc/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentSvcServer is the server API for DocumentSvc service.
// All implementations must embed UnimplementedDocumentSvcServer
// for forward compatibility
type DocumentSvcServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedDocumentSvcServer()
}

// UnimplementedDocumentSvcServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentSvcServer struct {
}

func (UnimplementedDocumentSvcServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedDocumentSvcServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedDocumentSvcServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDocumentSvcServer) mustEmbedUnimplementedDocumentSvcServer() {}

// UnsafeDocumentSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentSvcServer will
// result in compilation errors.
type UnsafeDocumentSvcServer interface {
	mustEmbedUnimplementedDocumentSvcServer()
}

func RegisterDocumentSvcServer(s grpc.ServiceRegistrar, srv DocumentSvcServer) {
	s.RegisterService(&DocumentSvc_ServiceDesc, srv)
}

func _DocumentSvc_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentSvcServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/documentsvc.DocumentSvc/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentSvcServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentSvc_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentSvcServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/documentsvc.DocumentSvc/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentSvcServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentSvc_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentSvcServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/documentsvc.DocumentSvc/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentSvcServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentSvc_ServiceDesc is the grpc.ServiceDesc for DocumentSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "documentsvc.DocumentSvc",
	HandlerType: (*DocumentSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _DocumentSvc_Hello_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _DocumentSvc_DeleteIndex_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _DocumentSvc_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "documentsvc.proto",
}