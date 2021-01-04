package ginutils

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	inTagPath     = "path"
	inTagQuery    = "query"
	inTagBody     = "body"
	inTagForm     = "form"
	inTagRequired = "required"

	tagNameIn = "in"
)

// Bind bind params from Path, Query, Body, Form. Donot support binary stream(files, images etc.)
// Support Tag `in`, specified that where we can get this value, only support one
// - path: from url path, don't support nested struct
// - query: from url query, don't support nested struct
// - body: from request's body, default use json, support nested struct
// - form: from request form
// - required: this value is not null
/*
type Example struct {
	ID   string `json:"id"   in:"path:id"`             // path value default is required
	Name string `json:"name" in:"query:name,required"` // query specified that get
}
*/
func Bind(c *gin.Context, param interface{}, decoders ...Decoder) (err error) {
	val := reflect.ValueOf(param)
	elm := reflect.Indirect(val)
	if val.Kind() != reflect.Ptr && elm.Kind() != reflect.Struct {
		err = fmt.Errorf("param must a pointer to struct, got %s", val.Kind().String())
		return
	}

	var decoder Decoder = new(jsonDecoder)
	if len(decoders) > 0 {
		decoder = decoders[0]
	}

	err = decoder.Decode(c.Request.Body, param)
	switch err {
	case io.EOF:
		err = nil
	default:
		err = fmt.Errorf("Decode body failed: %w", err)
		return
	}

	typ := elm.Type()
	for i := 0; i < elm.NumField(); i++ {
		field := elm.Field(i)
		fieldType := typ.Field(i)

		if isFromBody(fieldType.Tag, decoder.TagName()) {
			continue
		}

		// TODO
		field.Set(reflect.Value{})
	}

	return
}

func isFromBody(tag reflect.StructTag, bodyTag string) bool {
	bodyTag, bodyOK := tag.Lookup(bodyTag)
	if !bodyOK {
		return false
	}

	inTag, inOK := tag.Lookup(tagNameIn)
	// when we have a body tag, we don't need to specified in tag
	if !inOK {
		return true
	}

	// if we have body tag and in tag, check if in tag containt body
	return strings.Contains(inTag, inTagBody)
}

func isRequeired(tag string) bool {
	return strings.Contains(tag, inTagRequired)
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return false
}
