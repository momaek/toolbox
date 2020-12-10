package rpc

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Encoder encode http response body
type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

// Decoder decode http response body
type Decoder interface {
	Decode(io.Reader, interface{}) error
}

type coder interface {
	Encoder
	Decoder
}

type jsonCoder struct{}

// Encode implement Encoder
func (j *jsonCoder) Encode(iter interface{}) ([]byte, error) {
	return json.Marshal(iter)
}

// Decode implement Decoder
func (j *jsonCoder) Decode(r io.Reader, ret interface{}) error {
	return json.NewDecoder(r).Decode(ret)
}
