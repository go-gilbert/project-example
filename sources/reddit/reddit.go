package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-gilbert/project-example/server/feed"
)

type config struct {
	UserAgent  string   `json:"userAgent"`
	SubReddits []string `json:"subReddits"`
}

// NewDataSource creates a new data source
//
// This method called by server on extension load
func NewDataSource(rawCfg json.RawMessage) (feed.SourceReader, error) {
	cfg := new(config)
	if err := json.Unmarshal(rawCfg, cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration format (%s)", err)
	}

	return newDataSource(*cfg), nil
}
