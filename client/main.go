package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	documentsvc "github.com/frankcai/go-example/microservices/documentsvc/proto"
	ingestsvc "github.com/frankcai/go-example/microservices/ingestsvc/proto"
	"github.com/frankcai/go-example/microservices/proto/document"
	mono "github.com/frankcai/go-example/monolith/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/logger"
)

// set flags and default values
var microservices = flag.Bool("microservices", false, "call microservice or not")
var prefix = ""
var amountOfFiles = 2000
var allDocumentsSaved bool
var service micro.Service
var i ingestsvc.IngestSvcService
var d documentsvc.DocumentSvcService
var listeners []chan bool

// call like this "go run . -microservices=true" or "go run . -microservices=false"
func main() {
	flag.Parse()
	if *microservices {
		callMicroservices()
	} else {
		callMonolith()
	}
}

func callMonolith() {
	prefix = "Monolith/"

	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// shutdown client after processing finished
	go func() {
		<-time.After(time.Second * 5)
		cancel()
	}()

	// create a new service
	service = micro.NewService(
		micro.Name("client"),
		micro.AfterStart(func() error {
			// Use the generated client stub
			cl := mono.NewMonolithService("monolith", service.Client())

			// Wait until the monolith service is ready
			logger.Infof("Call %sHello", prefix)
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			for {
				heRsp, err := cl.Hello(ctx, &mono.HelloRequest{})
				if err != nil {
					fmt.Println(err)
				} else {
					logger.Infof("Call status: %t", heRsp.Succeeded)
					break
				}
			}

			// Delete existing bleve index from previous runs
			logger.Infof("Call %sDeleteIndex", prefix)
			ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			diRsp, err := cl.DeleteIndex(ctx, &mono.DeleteIndexRequest{
				IndexPath: "bleve",
			})
			if err != nil {
				fmt.Println(err)
				return nil
			}
			logger.Infof("Call status: %t", diRsp.Succeeded)

			// Create bleve index with fulltext parsed with apache tika
			logger.Infof("Call %sIngestFiles", prefix)
			ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			ciRsp, err := cl.IngestFiles(ctx, &mono.IngestFilesRequest{
				DataPath:  "data",
				IndexPath: "bleve",
			})
			if err != nil {
				fmt.Println(err)
				return nil
			}
			logger.Infof("Call status: %t", ciRsp.Succeeded)
			logger.Infof("Response: %s", ciRsp.CreateResponse)

			// Query bleve index for a specific string
			logger.Infof("Call %sQuery", prefix)
			defer cancel()
			queryRsp, err := cl.Query(ctx, &mono.QueryRequest{
				IndexPath:   "bleve",
				QueryString: "Statement",
			})
			if err != nil {
				fmt.Println(err)
				return nil
			}
			logger.Infof("Call status: %t", queryRsp.Succeeded)
			logger.Infof("Response: %s", queryRsp.QueryResponse)
			return nil
		}),
		micro.Context(ctx),
	)
	// start the service
	err := service.Run()
	if err != nil {
		logger.Fatal(err)
	}
}

// All methods of Sub will be executed when
// a message is received
type Sub struct{}

// Process events concerning the topic "DocumentSaved"
func (s *Sub) Process(ctx context.Context, event *document.DocumentSaved) error {
	// Run the query only if all documents have been sent
	amountOfFiles -= 1

	if amountOfFiles == 0 {
		logger.Infof("Document files have all been processed.")
		// Make request to query for a word
		logger.Infof("Call %sDocumentSvc/Query", prefix)
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
		dqRsp, err := d.Query(ctx, &documentsvc.QueryRequest{
			QueryString: "Statement",
		})
		if err != nil {
			fmt.Println(err)
		}
		logger.Infof("Call status: %t", dqRsp.Succeeded)
		logger.Infof("Response: %s", dqRsp.QueryResponse)
		SetAllDocumentsSaved(true)
		// Close all listeners:
		for _, listener := range listeners {
			close(listener)
		}
	}
	return nil
}

func GetChanADS() chan bool {
	listener := make(chan bool, 5)
	listeners = append(listeners, listener)
	return listener
}

func SetAllDocumentsSaved(newval bool) {
	allDocumentsSaved = newval
	for _, ch := range listeners {
		ch <- allDocumentsSaved
	}
}

func Background(name string, ch chan bool, done chan bool) {
	for v := range ch {
		logger.Infof("Value of allDocumentsSaved has changed to %t", v)
	}
	done <- false
}

func callMicroservices() {
	//log prefix
	prefix = "Microservices/"

	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// shutdown client after processing finished
	go func() {
		logger.Infof("open channel")
		l1 := GetChanADS()
		done := make(chan bool)
		// wait until
		go Background("B1", l1, done)
		<-done
		service.Options().Broker.Disconnect()
		cancel()
	}()

	// create a new service
	service = micro.NewService(
		micro.Name("client"),
		micro.Broker(broker.NewBroker(broker.Addrs(":33776"))),
		micro.Context(ctx),
		micro.AfterStart(func() error {
			// Use the generated client stub
			i = ingestsvc.NewIngestSvcService("ingestsvc", service.Client())
			d = documentsvc.NewDocumentSvcService("documentsvc", service.Client())

			// Wait until the ingest service is ready
			logger.Infof("Call %sIngestSvc/Hello", prefix)
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			for {
				heRsp, err := i.Hello(ctx, &ingestsvc.HelloRequest{})
				if err != nil {
					fmt.Println(err)
				} else {
					logger.Infof("Call status: %t", heRsp.Succeeded)
					break
				}
			}

			// Wait until the document service is ready
			logger.Infof("Call %sDocumentSvc/Hello", prefix)
			ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			for {
				heRsp, err := d.Hello(ctx, &documentsvc.HelloRequest{})
				if err != nil {
					fmt.Println(err)
				} else {
					logger.Infof("Call status: %t", heRsp.Succeeded)
					break
				}
			}

			// Delete existing bleve index from previous runs
			logger.Infof("Call %sDocumentSvc/DeleteIndex", prefix)
			ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			diRsp, err := d.DeleteIndex(ctx, &documentsvc.DeleteIndexRequest{
				IndexPath: "bleve",
			})
			if err != nil {
				fmt.Println(err)
				return nil
			}
			logger.Infof("Call status: %t", diRsp.Succeeded)

			// Make request to ingest files
			logger.Infof("Call %sIngestSvc/IngestFiles", prefix)
			ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
			defer cancel()
			ifRsp, err := i.IngestFiles(ctx, &ingestsvc.IngestFilesRequest{
				DataPath: "data",
			})
			if err != nil {
				fmt.Println(err)
			}
			logger.Infof("Call status: %t", ifRsp.Succeeded)
			logger.Infof("Response: %s", ifRsp.IngestResponse)
			return nil
		}),
	)

	// register subscriber
	micro.RegisterSubscriber("DocumentSaved", service.Server(), new(Sub))

	// start the service to wait as a subscriber
	err := service.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
