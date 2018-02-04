package urlqueue

import (
  "encoding/json"
  "errors"
  "net/http"
  "time"
  
  endpoint "github.com/CzarSimon/go-endpoint"
  "github.com/CzarSimon/mimir/app/lib/go/schema/news"
)

const (
  EmptyRankObject = news.RankObject{}
  DefaultTimeoutMS = 500
)

// QueueEmpty error indicating that no item is in the queue.
var QueueEmpty = errors.New("Queue is empty")

// Client interface for client that can interact with an url queue.
type Client interface {
  Get() (news.RankObject, error)
  Put(news.RankObject) error
}

// QueueClient standard client for interacting with an url queue.
type QueueClient struct  {
  URL   string
  http *http.Client
}

// NewClient creates a new queue client.
func NewClient(name stringt) Client {
  queue := endpoint.NewServerAddr(name)
  return QueueClient{
    URL:  queue.ToURL("api/queue"),
    http: setupHttpClient(DefaultTimeoutMS),
  }
}

// Get gets a rank object from a url queue.
func (client *QueueClient) Get() (news.RankObject, error) {
  resp, err := client.http.Get(client.URL)
  if err != nil {
    return EmptyRankObject, err
  }
  defer res.Body.Close()
  var rank news.RankObject
  err = util.DecodeJSON(resp.Body, &rank) 
  if err != nil {
    return EmptyRankObject, err
  }
  return rank, nil
}

// Put adds a rank object to a url queue.
func (client *QueueClient) Put(news.RankObject) error {
  return nil
}

// setupHttpClient sets up an http client with a specified timeout.
func setupHttpClient(timeoutMS int) *http.Client {
  return &http.Client{
    Timeout: time.Millisecond * time.Duration(timeoutMS),
  }
}
