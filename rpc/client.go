package client

import (
	"net/http"

	"github.com/momaek/toolbox/logger"
)

// Client ..
type Client struct {
	*http.Client
}

// New with default http client
func New() *Client {
	return &Client{Client: http.DefaultClient}
}

// NewWithRoundTripper new with roundtripper
func NewWithRoundTripper(tr http.RoundTripper) *Client {
	return &Client{Client: &http.Client{Transport: tr}}
}

// NewWithHTTPClient new with http client
func NewWithHTTPClient(c *http.Client) *Client {
	return &Client{Client: c}
}

// Get ..
func (c *Client) Get(l logger.Logger, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	return c.do(l, req)
}

// Head ..
func (c *Client) Head(l logger.Logger, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return
	}

	return c.do(l, req)
}
