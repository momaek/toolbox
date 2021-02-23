package ginutils

import (
	"encoding/json"
	"io"
)

// Decoder decoder..
type Decoder interface {
	// if empty reader should return io.EOF error
	Decode(io.Reader, interface{}) error
	TagName() string
}

type jsonDecoder struct{}

// Decode implement Decoder
func (*jsonDecoder) Decode(r io.Reader, p interface{}) error {
	return json.NewDecoder(r).Decode(p)
}

// TagName return support struct tag name
func (*jsonDecoder) TagName() string { return "json" }
