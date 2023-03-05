package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	proto "github.com/frankcai/go-example/microservices/ingestsvc/proto"
	"github.com/frankcai/go-example/microservices/proto/document"
	tika "github.com/google/go-tika/tika"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/logger"
)

// set flags and default values
var dataDir = flag.String("datadir", "data/", "data directory")
var indexPath = flag.String("index", "docs.bleve", "index path")
var pub micro.Event

// send events using the publisher
func sendEv(p micro.Publisher, ev *document.NewDocument) {
	// publish an event
	if err := p.Publish(context.Background(), ev); err != nil {
		logger.Infof("error publishing: %v", err)
	}
}

type IngestSvc struct{}

// RPC function for returning true if service active
func (g *IngestSvc) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	logger.Infof("Returning Hello from IngestSvc")
	rsp.Succeeded = true
	return nil
}

// RPC function for ingesting files
func (g *IngestSvc) IngestFiles(ctx context.Context, req *proto.IngestFilesRequest, rsp *proto.IngestFilesResponse) error {
	// setup apache tika
	tikaServer := startTika()
	defer tikaServer.Stop()
	tikaClient := tika.NewClient(nil, tikaServer.URL())
	var wg sync.WaitGroup

	// open the directory
	dirEntries, err := os.ReadDir(*dataDir)
	if err != nil {
		return err
	}

	// walk the directory entries for indexing
	logger.Infof("Indexing...")
	count := 0
	startTime := time.Now()

	for _, dirEntry := range dirEntries {
		filename := dirEntry.Name()
		// parse as text
		body := tikaParse(*tikaClient, *dataDir+filename)
		if err != nil {
			return err
		}

		ext := filepath.Ext(filename)
		docID := filename[:(len(filename) - len(ext))]

		// create new event
		ev := &document.NewDocument{
			Id:        docID,
			Timestamp: time.Now().Unix(),
			Filename:  filename,
			Ext:       ext,
			Message:   body,
		}
		// Increment the WaitGroup counter
		wg.Add(1)
		go func() {
			// Pub to topic NewDocument
			sendEv(pub, ev)
			// Decrement the WaitGroup counter
			defer wg.Done()
		}()
		count++
	}

	indexDuration := time.Since(startTime)
	indexDurationSeconds := float64(indexDuration) / float64(time.Second)
	timePerDoc := float64(indexDuration) / float64(count)
	var resulttimes = "Indexed " + strconv.Itoa(count) + " documents, in " + strconv.FormatFloat(indexDurationSeconds, 'f', 2, 64) + "s (average " + strconv.FormatFloat(timePerDoc/float64(time.Millisecond), 'f', 2, 64) + "ms/doc)"
	logger.Infof(resulttimes)
	rsp.Succeeded = true
	rsp.IngestResponse = resulttimes
	wg.Wait()
	return nil
}

func main() {
	flag.Parse()
	logger.Infof("Data path: %s", *dataDir)
	logger.Infof("Index path: %s", *indexPath)

	// create a new service
	service := micro.NewService(
		micro.Name("ingestsvc"),
		micro.Address(":6778"),
		micro.Broker(broker.NewBroker(broker.Addrs(":33778"))),
	)

	// initialise flags
	service.Init()

	proto.RegisterIngestSvcHandler(service.Server(), new(IngestSvc))

	// create publisher
	pub = micro.NewEvent("NewDocument", service.Client())

	// start the service
	err := service.Run()
	if err != nil {
		logger.Fatal(err)
	}
}

// Start apache tika and return tika client
func startTika() (s *tika.Server) {
	s, err := tika.NewServer("tika-server-1.21.jar", "")
	if err != nil {
		logger.Fatal(err)
	}
	err = s.Start(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	return s
}

// Parses a given file with apache tika
func tikaParse(c tika.Client, file string) (body string) {
	f, err := os.Open(file)
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()
	body, err = c.Parse(context.Background(), f)
	if err != nil {
		logger.Fatal(err)
	}
	return
}
