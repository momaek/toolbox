package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	json "github.com/json-iterator/go"
	"github.com/momaek/toolbox/logger"
)

func (c *Client) do(l logger.Logger, req *http.Request) (resp *http.Response, err error) {
	var (
		url    = req.URL.String()
		method = req.Method
	)

	l.WithField(map[string]interface{}{
		"type":   "RPC",
		"url":    url,
		"method": method,
	}).Info("[Started]")

	now := time.Now()
	resp, err = c.Client.Do(req)
	l.WithField(map[string]interface{}{
		"type":    "RPC",
		"url":     url,
		"method":  method,
		"status":  resp.StatusCode,
		"latency": time.Since(now),
	}).Info("[Completed]")

	return
}

func (c *Client) doRet(l logger.Logger, req *http.Request, ret interface{}) (err error) {
	return
}

func callRet(resp *http.Response, ret interface{}) (err error) {
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength > 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
		}

		return
	}

	e := &Error{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}

	if resp.ContentLength > 0 {
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, resp.Body)
		e.Body = buf.String()
	}

	return e
}
