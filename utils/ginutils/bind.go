package ginutils

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	inTagPath     = "path"
	inTagQuery    = "query"
	inTagBody     = "body"
	inTagForm     = "form"
	inTagRequired = "required"
	inTagNotNull  = "notnull"

	tagNameIn = "in"
	tagSep    = ","
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
	if err != nil {
		switch err {
		case io.EOF:
			err = nil
		default:
			err = fmt.Errorf("Decode body failed: %w", err)
			return
		}
	}

	typ := elm.Type()
	for i := 0; i < elm.NumField(); i++ {
		field := elm.Field(i)
		fieldType := typ.Field(i)

		if isFromBody(fieldType.Tag, decoder.TagName()) {
			continue
		}

		inTag := fieldType.Tag.Get(tagNameIn)
		if len(inTag) == 0 {
			continue
		}

		var (
			loc, name   = getInTagParamLocAndName(inTag)
			val         string
			isRequeired = isRequeired(inTag)
			ok          bool
		)

		switch loc {
		case inTagPath:
			val = c.Param(name)
			if isRequeired && len(val) == 0 {
				err = fmt.Errorf("%s is required", name)
				return
			}
			ok = true
		case inTagForm:
			val, ok = c.GetPostForm(name)
		case inTagQuery:
			val, ok = c.GetQuery(name)
		default:
			err = fmt.Errorf("unsupportted location tag: %s", loc)
			return
		}

		if isRequeired && !ok {
			err = fmt.Errorf("%s is required", name)
			return
		}

		if isNotNull(inTag) && len(val) == 0 {
			err = fmt.Errorf("%s can't be empty", name)
			return
		}

		reflectVal := bind(val, field.Type())
		if reflectVal.Type().ConvertibleTo(field.Type()) {
			field.Set(reflectVal)
		}
	}

	return
}

// getInTagParamLoc `in:"query:xxx,xxxxxxxxx"`
func getInTagParamLocAndName(tag string) (loc, name string) {
	splits := strings.Split(tag, tagSep)
	locs := strings.Split(splits[0], ":")
	if len(locs) != 2 {
		return
	}

	loc = locs[0]
	name = locs[1]
	return
}

func isFromBody(tag reflect.StructTag, tagName string) bool {
	_, bodyOK := tag.Lookup(tagName)
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

func isNotNull(tag string) bool {
	return strings.Contains(tag, inTagNotNull)
}

func stringBinder(val string, typ reflect.Type) reflect.Value {
	return reflect.ValueOf(val)
}

func uintBinder(val string, typ reflect.Type) reflect.Value {
	if len(val) == 0 {
		return reflect.Zero(typ)
	}

	uintValue, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return reflect.Zero(typ)
	}
	pValue := reflect.New(typ)
	pValue.Elem().SetUint(uintValue)
	return pValue.Elem()
}

func intBinder(val string, typ reflect.Type) reflect.Value {
	if len(val) == 0 {
		return reflect.Zero(typ)
	}
	intValue, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return reflect.Zero(typ)
	}
	pValue := reflect.New(typ)
	pValue.Elem().SetInt(intValue)
	return pValue.Elem()
}

func floatBinder(val string, typ reflect.Type) reflect.Value {
	if len(val) == 0 {
		return reflect.Zero(typ)
	}
	floatValue, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return reflect.Zero(typ)
	}
	pValue := reflect.New(typ)
	pValue.Elem().SetFloat(floatValue)
	return pValue.Elem()
}

func boolBinder(val string, typ reflect.Type) reflect.Value {
	v := strings.TrimSpace(strings.ToLower(val))
	switch v {
	case "true":
		return reflect.ValueOf(true)
	}
	// Return false by default.
	return reflect.ValueOf(false)
}

func timeBinder(val string, typ reflect.Type) reflect.Value {
	for _, f := range TimeFormats {
		if f == "" {
			continue
		}

		if strings.Contains(f, "07") || strings.Contains(f, "MST") {
			if r, err := time.Parse(f, val); err == nil {
				return reflect.ValueOf(r)
			}
		} else {
			if r, err := time.ParseInLocation(f, val, time.Local); err == nil {
				return reflect.ValueOf(r)
			}
		}
	}

	if unixInt, err := strconv.ParseInt(val, 10, 64); err == nil {
		return reflect.ValueOf(time.Unix(unixInt, 0))
	}

	return reflect.Zero(typ)
}

func pointerBinder(val string, typ reflect.Type) reflect.Value {
	if len(val) == 0 {
		return reflect.Zero(typ)
	}

	v := bind(val, typ.Elem())
	p := reflect.New(v.Type()).Elem()
	p.Set(v)
	return p.Addr()
}

const (
	// DefaultDateFormat day
	DefaultDateFormat = "2006-01-02"
	// DefaultDatetimeFormat minute
	DefaultDatetimeFormat = "2006-01-02 15:0"
	// DefaultDatetimeFormatSecond second
	DefaultDatetimeFormatSecond = "2006-01-02 15:04:05"
)

func bind(val string, typ reflect.Type) reflect.Value {
	binder, ok := TypeBinders[typ]
	if !ok {
		binder, ok = KindBinders[typ.Kind()]
		if !ok {
			// WARN.Println("no binder for type:", typ)
			// TODO slice | struct
			return reflect.Zero(typ)
		}
	}

	return binder(val, typ)
}

type binder func(string, reflect.Type) reflect.Value

var (
	// TimeFormats supported time formats, also support unix time and time.RFC3339.
	TimeFormats []string

	// TypeBinders bind type
	TypeBinders = make(map[reflect.Type]binder)

	// KindBinders bind kind
	KindBinders = make(map[reflect.Kind]binder)
)

func init() {
	KindBinders[reflect.Int] = intBinder
	KindBinders[reflect.Int8] = intBinder
	KindBinders[reflect.Int16] = intBinder
	KindBinders[reflect.Int32] = intBinder
	KindBinders[reflect.Int64] = intBinder

	KindBinders[reflect.Uint] = uintBinder
	KindBinders[reflect.Uint8] = uintBinder
	KindBinders[reflect.Uint16] = uintBinder
	KindBinders[reflect.Uint32] = uintBinder
	KindBinders[reflect.Uint64] = uintBinder

	KindBinders[reflect.Float32] = floatBinder
	KindBinders[reflect.Float64] = floatBinder

	KindBinders[reflect.String] = stringBinder
	KindBinders[reflect.Bool] = boolBinder
	KindBinders[reflect.Ptr] = pointerBinder

	TypeBinders[reflect.TypeOf(time.Time{})] = timeBinder

	TimeFormats = append(TimeFormats, DefaultDateFormat, DefaultDatetimeFormat, DefaultDatetimeFormatSecond, time.RFC3339)
}
