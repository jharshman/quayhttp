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

func (c *Client)HTTPClient() *http.Client {
  if c.HTTPClient == nil {
    c.HTTPClient = &http.Client{}
  }
}

func (c *Client)NewRequest(ctx, method, requestPath string) *http.Request {
  // create the request to send
  req, _ := http.NewRequest(method, c.Address(), nil)
  req = req.WithContext(ctx)

  return req
}

func (c *Client)DoRequest() {
  // send the request
}
