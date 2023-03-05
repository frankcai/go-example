package main

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"time"

	bleve "github.com/blevesearch/bleve/v2"
	tika "github.com/google/go-tika/tika"
	"go-micro.dev/v4/logger"
)

// Start apache tika and return tika client
func startTika() (s *tika.Server) {
	s, err := tika.NewServer("tika-server-1.21.jar", "9998")
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
		logger.Fatal(err, " ", file)
	}
	return
}

// Index all files in jsonDir into bleve index
func indexFiles(i bleve.Index, t *tika.Client, d string) (string, error) {
	// print current directory
	path, err := os.Getwd()
	if err != nil {
		logger.Infof(err.Error())
	}
	logger.Infof("Current path: " + path)

	// open the directory
	dirEntries, err := os.ReadDir(d)
	if err != nil {
		return "", err
	}

	// walk the directory entries for indexing
	logger.Infof("Indexing...")
	count := 0
	startTime := time.Now()
	for _, dirEntry := range dirEntries {
		filename := dirEntry.Name()
		// parse as text
		absDatadir, _ := filepath.Abs(d)
		body := tikaParse(*t, absDatadir+"/"+filename)
		if err != nil {
			return "", err
		}

		ext := filepath.Ext(filename)
		docID := filename[:(len(filename) - len(ext))]
		i.Index(docID, body)
		count++
	}
	indexDuration := time.Since(startTime)
	indexDurationSeconds := float64(indexDuration) / float64(time.Second)
	timePerDoc := float64(indexDuration) / float64(count)
	var resulttimes = "Indexed " + strconv.Itoa(count) + " documents, in " + strconv.FormatFloat(indexDurationSeconds, 'f', 2, 64) + "s (average " + strconv.FormatFloat(timePerDoc/float64(time.Millisecond), 'f', 2, 64) + "ms/doc)"
	logger.Infof(resulttimes)
	return resulttimes, nil
}
