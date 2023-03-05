package main

import (
	"context"
	"flag"
	"os"
	"time"

	bleve "github.com/blevesearch/bleve/v2"
	proto "github.com/frankcai/go-example/microservices/documentsvc/proto"
	"github.com/frankcai/go-example/microservices/proto/document"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/logger"
)

// set flags and default values
var indexPath = flag.String("index", "bleve", "index path")
var fileCounter = 0
var bleveIndex bleve.Index
var pub micro.Event

// send events using the publisher
func sendEv(p micro.Publisher, ev *document.DocumentSaved) {
	// publish an event
	if err := p.Publish(context.Background(), ev); err != nil {
		logger.Infof("error publishing: %v", err)
	}
}

// All methods of Sub will be executed when
// a message is received
type Sub struct{}

// Process events concerning the topic "NewDocument"
func (s *Sub) Process(ctx context.Context, event *document.NewDocument) error {
	fileCounter++
	// index into bleve
	err := indexFiles(bleveIndex, event.Id, event.Message)
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	// create new event
	ev := &document.DocumentSaved{
		Id:        event.Id,
		Timestamp: time.Now().Unix(),
		Filename:  event.Filename,
		Ext:       event.Ext,
	}
	// pub to topic
	go sendEv(pub, ev)
	return nil
}

func indexFiles(bleveIndex bleve.Index, id string, body string) error {
	return bleveIndex.Index(id, body)
}

type DocumentSvc struct{}

// RPC function for returning true if service active
func (g *DocumentSvc) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	logger.Infof("Returning Hello from DocumentSvc")
	rsp.Succeeded = true
	return nil
}

// RPC function for deleting bleve index
func (g *DocumentSvc) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest, rsp *proto.DeleteIndexResponse) error {
	logger.Infof("Delete index at path: %s", req.IndexPath)
	err := os.RemoveAll(req.IndexPath)
	if err != nil {
		rsp.Succeeded = false
	} else {
		rsp.Succeeded = true
		// create new index
		createIndex()
	}

	return nil
}

// RPC function for querying an index
func (g *DocumentSvc) Query(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryResponse) error {
	rsp.Succeeded = true
	rsp.QueryResponse = queryText(bleveIndex, req.QueryString)
	return nil
}

func createIndex() {
	// initialise bleve index
	logger.Infof("Try opening existing index at path: %s", *indexPath)
	var err error
	bleveIndex, err = bleve.Open(*indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		logger.Error("Index doesn't exist")
		logger.Infof("Creating new index")
		logger.Infof("1/2 Build mapping")
		indexMapping := bleve.NewIndexMapping()
		logger.Infof("2/2 Create index")
		bleveIndex, err = bleve.New(*indexPath, indexMapping)
		if err != nil {
			logger.Fatal(err)
		}
	} else if err != nil {
		logger.Fatal(err)
	} else {
		logger.Infof("Opening existing index")
	}
	// skip unused variable error
	_ = bleveIndex
}

func main() {

	// create a new service
	service := micro.NewService(
		micro.Name("documentsvc"),
		micro.Address(":6779"),
		micro.Broker(broker.NewBroker(broker.Addrs(":33779"))),
		micro.AfterStop(Close),
	)

	// initialise flags
	service.Init()

	proto.RegisterDocumentSvcHandler(service.Server(), new(DocumentSvc))

	// register subscriber
	micro.RegisterSubscriber("NewDocument", service.Server(), new(Sub))
	// create publisher
	pub = micro.NewEvent("DocumentSaved", service.Client())

	// start the service
	err := service.Run()
	if err != nil {
		logger.Fatal(err)
	}
}

func Close() (err error) {
	bleveIndex.Close()
	return
}
