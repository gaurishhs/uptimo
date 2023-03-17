package fetcher

import (
	"net/http"
	"time"
)

type Fetcher struct {
  client *http.Client
}

type FetcherConfig struct {
  // Timeout is the timeout for the Fetcher 
  Timeout time.Duration
  // Proxy is the proxy to use for the Fetcher 
  Proxy string 
  // MaxIdleConns is the maximum number of idle connections to keep 
  MaxIdleConns int 
}

// NewFetcher creates a new fetcher
func NewFetcher(config *FetcherConfig) *Fetcher {
  // Create a new http client
  client := &http.Client{
    Timeout: config.Timeout,
    Transport: &http.Transport{
      Proxy: http.ProxyURL(config.Proxy), 
      MaxIdleConns: config.MaxIdleConns,
    },
  }

  return &Fetcher{
    client: client,
  }
}

// Make the request
func (c *Fetcher) Request() {}
