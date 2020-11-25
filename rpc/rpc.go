package rpc

import "net/http"

var UserAgent = "github.com/momaek/toolbox/rpc"

const defaultTimestampFormat = "2006-01-02 15:04:05.999999"

// Client client
type Client struct {
	*http.Client
}
