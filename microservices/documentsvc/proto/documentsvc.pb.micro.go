// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: documentsvc.proto

package documentsvc

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DocumentSvc service

func NewDocumentSvcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DocumentSvc service

type DocumentSvcService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...client.CallOption) (*DeleteIndexResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
}

type documentSvcService struct {
	c    client.Client
	name string
}

func NewDocumentSvcService(name string, c client.Client) DocumentSvcService {
	return &documentSvcService{
		c:    c,
		name: name,
	}
}

func (c *documentSvcService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "DocumentSvc.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentSvcService) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...client.CallOption) (*DeleteIndexResponse, error) {
	req := c.c.NewRequest(c.name, "DocumentSvc.DeleteIndex", in)
	out := new(DeleteIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentSvcService) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.name, "DocumentSvc.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DocumentSvc service

type DocumentSvcHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
	DeleteIndex(context.Context, *DeleteIndexRequest, *DeleteIndexResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
}

func RegisterDocumentSvcHandler(s server.Server, hdlr DocumentSvcHandler, opts ...server.HandlerOption) error {
	type documentSvc interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
		DeleteIndex(ctx context.Context, in *DeleteIndexRequest, out *DeleteIndexResponse) error
		Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error
	}
	type DocumentSvc struct {
		documentSvc
	}
	h := &documentSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&DocumentSvc{h}, opts...))
}

type documentSvcHandler struct {
	DocumentSvcHandler
}

func (h *documentSvcHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.DocumentSvcHandler.Hello(ctx, in, out)
}

func (h *documentSvcHandler) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, out *DeleteIndexResponse) error {
	return h.DocumentSvcHandler.DeleteIndex(ctx, in, out)
}

func (h *documentSvcHandler) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.DocumentSvcHandler.Query(ctx, in, out)
}
