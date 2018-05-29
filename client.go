package quayhttp

import (
  "net/http"
)

type Client struct {
  HTTPClient *http.Client
  Address string
  Token string
}

func (c *Client)Address() string {
  var addr string
  if c.Address == "" {
    addr = "quay.io"
  }
  return addr
}

func (c *Client)NewRequest(method, requestPath string) *http.Request {
  // create the request to send
}

func (c *Client)DoRequest() {
  // send the request
}
