package rpc

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/momaek/toolbox/logger"
	"github.com/momaek/toolbox/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

	return callRet(l, resp, ret)
}

// CallRet ...
func CallRet(l logger.Logger, resp *http.Response, ret interface{}) (err error) {
	return callRet(l, resp, ret)
}

func callRet(l logger.Logger, resp *http.Response, ret interface{}) (err error) {
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
