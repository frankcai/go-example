// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: monolith.proto

package monolith

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

// Api Endpoints for Monolith service

func NewMonolithEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Monolith service

type MonolithService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
	IngestFiles(ctx context.Context, in *IngestFilesRequest, opts ...client.CallOption) (*IngestFilesResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...client.CallOption) (*DeleteIndexResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
}

type monolithService struct {
	c    client.Client
	name string
}

func NewMonolithService(name string, c client.Client) MonolithService {
	return &monolithService{
		c:    c,
		name: name,
	}
}

func (c *monolithService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "Monolith.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monolithService) IngestFiles(ctx context.Context, in *IngestFilesRequest, opts ...client.CallOption) (*IngestFilesResponse, error) {
	req := c.c.NewRequest(c.name, "Monolith.IngestFiles", in)
	out := new(IngestFilesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monolithService) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...client.CallOption) (*DeleteIndexResponse, error) {
	req := c.c.NewRequest(c.name, "Monolith.DeleteIndex", in)
	out := new(DeleteIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monolithService) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.name, "Monolith.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monolith service

type MonolithHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
	IngestFiles(context.Context, *IngestFilesRequest, *IngestFilesResponse) error
	DeleteIndex(context.Context, *DeleteIndexRequest, *DeleteIndexResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
}

func RegisterMonolithHandler(s server.Server, hdlr MonolithHandler, opts ...server.HandlerOption) error {
	type monolith interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
		IngestFiles(ctx context.Context, in *IngestFilesRequest, out *IngestFilesResponse) error
		DeleteIndex(ctx context.Context, in *DeleteIndexRequest, out *DeleteIndexResponse) error
		Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error
	}
	type Monolith struct {
		monolith
	}
	h := &monolithHandler{hdlr}
	return s.Handle(s.NewHandler(&Monolith{h}, opts...))
}

type monolithHandler struct {
	MonolithHandler
}

func (h *monolithHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.MonolithHandler.Hello(ctx, in, out)
}

func (h *monolithHandler) IngestFiles(ctx context.Context, in *IngestFilesRequest, out *IngestFilesResponse) error {
	return h.MonolithHandler.IngestFiles(ctx, in, out)
}

func (h *monolithHandler) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, out *DeleteIndexResponse) error {
	return h.MonolithHandler.DeleteIndex(ctx, in, out)
}

func (h *monolithHandler) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.MonolithHandler.Query(ctx, in, out)
}
