package main

import (
	"context"
	"flag"
	"os"

	bleve "github.com/blevesearch/bleve/v2"
	proto "github.com/frankcai/go-example/monolith/proto"
	tika "github.com/google/go-tika/tika"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/logger"
)

// set flags and default values
var dataDir = flag.String("datadir", "data", "data directory")
var indexPath = flag.String("index", "bleve", "path of the bleve index")
var tikaServer tika.Server

// Creates or opens an existing bleve index
func main() {
	flag.Parse()
	logger.Infof("Data path: %s", *dataDir)
	logger.Infof("Index path: %s", *indexPath)

	// create a new service
	service := micro.NewService(
		micro.Name("monolith"),
		micro.Address(":6777"),
		micro.Broker(broker.NewBroker(broker.Addrs(":33777"))),
	)

	// initialise flags
	service.Init()

	proto.RegisterMonolithHandler(service.Server(), new(Monolith))

	// start the service
	err := service.Run()
	if err != nil {
		logger.Fatal(err)
	}
}

type Monolith struct{}

// RPC function for returning true if service active
func (g *Monolith) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	logger.Infof("Returning Hello from Monolith")
	rsp.Succeeded = true
	return nil
}

// RPC function for ingesting files
func (g *Monolith) IngestFiles(ctx context.Context, req *proto.IngestFilesRequest, rsp *proto.IngestFilesResponse) error {
	var createRsp string
	// start apache tika
	tikaServer := startTika()
	logger.Infof("Running Apache Tika at: %s", tikaServer.URL())
	defer tikaServer.Stop()
	tikaClient := tika.NewClient(nil, tikaServer.URL())

	logger.Infof("Try opening existing index at path: %s", req.IndexPath)
	index, err := bleve.Open(req.IndexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		logger.Error("Index doesn't exist")
		logger.Infof("Creating new index")
		logger.Infof("1/3 Build mapping")
		indexMapping := bleve.NewIndexMapping()
		logger.Infof("2/3 Create index")
		index, err = bleve.New(req.IndexPath, indexMapping)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Infof("3/3 Index files")
		createRsp, err = indexFiles(index, tikaClient, req.DataPath)
		if err != nil {
			logger.Fatal(err)
		}
	} else if err != nil {
		logger.Fatal(err)
	} else {
		logger.Infof("Opening existing index")
	}
	defer index.Close()
	rsp.Succeeded = true
	rsp.CreateResponse = createRsp
	return nil
}

// RPC function for querying an index
func (g *Monolith) Query(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryResponse) error {
	logger.Infof("Open index for query at path: %s", req.IndexPath)
	// open the index
	index, err := bleve.Open(req.IndexPath)
	if err != nil {
		logger.Error("Failure opening index")
		logger.Fatal(err)
		rsp.Succeeded = false
	} else {
		logger.Infof("Index opened succesfully")
		rsp.Succeeded = true
		rsp.QueryResponse = queryText(index, req.QueryString)
	}
	defer index.Close()
	return nil
}

// RPC function for deleting an index
func (g *Monolith) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest, rsp *proto.DeleteIndexResponse) error {
	logger.Infof("Delete index at path: %s", req.IndexPath)
	err := os.RemoveAll(req.IndexPath)
	if err != nil {
		rsp.Succeeded = false
	} else {
		rsp.Succeeded = true
	}

	return nil
}
