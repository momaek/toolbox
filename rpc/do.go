package rpc

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/momaek/toolbox/logger"
	"github.com/momaek/toolbox/utils"
)

func (c *Client) do(l logger.Logger, req *http.Request) (resp *http.Response, err error) {
	var (
		url    = req.URL.String()
		method = req.Method
	)

	logger.NewWithoutCaller(l.ReqID()).WithField(map[string]interface{}{
		"type":   "RPC",
		"url":    url,
		"method": method,
	}).Info("[Started]")

	now := time.Now()
	resp, err = c.Client.Do(req)
	logger.NewWithoutCaller(l.ReqID()).WithField(map[string]interface{}{
		"type":    "RPC",
		"url":     url,
		"method":  method,
		"status":  resp.StatusCode,
		"latency": time.Since(now).String(),
	}).Info("[Completed]")

	return
}

func (c *Client) doRet(l logger.Logger, req *http.Request, ret interface{}) (err error) {
	resp, err := c.do(l, req)
	if err != nil {
		return
	}

	return c.callRet(l, resp, ret)
}

// CallRet ...
func CallRet(l logger.Logger, resp *http.Response, ret interface{}) (err error) {
	return New().callRet(l, resp, ret)
}

func (c *Client) callRet(l logger.Logger, resp *http.Response, ret interface{}) (err error) {
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength > 0 {
			err = c.coder.Decode(resp.Body, ret)
			return
		}
	}

	e := &Error{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}

	if xlog := resp.Header.Get(logger.XLogKey); len(xlog) > 0 {
		l.Xput([]string{xlog})
		e.Detail = xlog
	}

	if resp.ContentLength > 0 {
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, resp.Body)
		e.Body = utils.BytesToString(buf.Bytes())
	}

	return e
}
