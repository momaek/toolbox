package rpc

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"

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

// Post ...
func (c *Client) Post(l logger.Logger, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return
	}

	return c.do(l, req)
}

// Delete delete method
func (c *Client) Delete(l logger.Logger, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}

	return c.do(l, req)
}

// Patch http patch method
func (c *Client) Patch(l logger.Logger, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return
	}

	return c.do(l, req)
}

// Put put
func (c *Client) Put(l logger.Logger, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return
	}

	return c.do(l, req)
}

// PostEx post with body
func (c *Client) PostEx(l logger.Logger, url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return
	}

	return c.do(l, req)
}

// PostWith ..
func (c *Client) PostWith(l logger.Logger, url string, body io.Reader, bodyType string, bodyLength int64) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	req.ContentLength = bodyLength
	return c.do(l, req)
}

// PostForm post content type form
func (c *Client) PostForm(l logger.Logger, url1 string, data map[string][]string) (*http.Response, error) {
	msg := url.Values(data).Encode()
	return c.PostWith(l, url1, strings.NewReader(msg), "application/x-www-form-urlencoded", int64(len(msg)))
}

// PostJSON ..
func (c *Client) PostJSON(l logger.Logger, url1 string, data interface{}) (*http.Response, error) {
	var (
		b   []byte
		err error
	)
	if data != nil {
		b, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	return c.PostWith(l, url1, bytes.NewBuffer(b), "application/json", int64(len(b)))
}

// Call make a post request and parse json response, content-type: "application/x-www-form-urlencoded"
func (c *Client) Call(l logger.Logger, url1 string, ret interface{}) (err error) {
	resp, err := c.PostForm(l, url1, nil)
	if err != nil {
		return
	}

	return callRet(l, resp, ret)
}

// CallWith call with specified method,url,body... and parse json request
func (c *Client) CallWith(l logger.Logger, method, url1, contentType string, body io.Reader, bodyLength int64, ret interface{}) (err error) {
	req, err := http.NewRequest(method, url1, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)
	req.ContentLength = bodyLength
	return c.doRet(l, req, ret)
}

// CallWithJSON call with json
func (c *Client) CallWithJSON(l logger.Logger, url1 string, data interface{}, ret interface{}) (err error) {
	var b []byte
	if data != nil {
		b, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	return c.CallWith(l, http.MethodPost, url1, "application/json", bytes.NewBuffer(b), int64(len(b)), ret)
}

// CallWithForm call with content-type "application/x-www-form-urlencoded"
func (c *Client) CallWithForm(l logger.Logger, url1 string, data map[string][]string, ret interface{}) (err error) {
	msg := url.Values(data).Encode()
	return c.CallWith(l, http.MethodPost, url1, "application/x-www-form-urlencoded", strings.NewReader(msg), int64(len(msg)), ret)
}

// GetCall ..
func (c *Client) GetCall(l logger.Logger, url string, ret interface{}) (err error) {
	resp, err := c.Get(l, url)
	if err != nil {
		return
	}

	return callRet(l, resp, ret)
}

// DeleteCall ..
func (c *Client) DeleteCall(l logger.Logger, url string, ret interface{}) (err error) {
	resp, err := c.Delete(l, url)
	if err != nil {
		return
	}

	return callRet(l, resp, ret)
}

// PutCall ..
func (c *Client) PutCall(l logger.Logger, url string, ret interface{}) (err error) {
	resp, err := c.Put(l, url)
	if err != nil {
		return
	}

	return callRet(l, resp, ret)
}
