package main

import (
	bleve "github.com/blevesearch/bleve/v2"
	"go-micro.dev/v4/logger"
)

// Queries a bleve index for a specific query string
func queryText(i bleve.Index, querystring string) (response string) {
	logger.Infof("Searching for '%s'", querystring)
	// search for some text
	query := bleve.NewMatchQuery(querystring)
	search := bleve.NewSearchRequest(query)
	searchResults, err := i.Search(search)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof(searchResults.String())
	response = searchResults.String()
	return
}
