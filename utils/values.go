package utils

import (
	"reflect"
	"strconv"
)

// Value 帮助你从任意 interface{} 转换到各种值类型
// Must* 系列返回的值，源类型复合的情况下才返回
// 其他返回错误的系列，是强转换
type Value struct {
	value interface{}
}

// Bool interface{} 转成 bool 类型
func (v *Value) Bool() (bool, error) {
	if ret, ok := v.value.(bool); ok {
		return ret, nil
	}
	if v.value == "on" {
		return true, nil
	}
	return strconv.ParseBool(v.String())
}

// Float32 interface{} 转成 float32 类型
func (v *Value) Float32() (float32, error) {
	if ret, ok := v.value.(float32); ok {
		return ret, nil
	}
	value, err := strconv.ParseFloat(v.String(), 32)
	return float32(value), err
}

// Float64 interface{} 转成 Float64 类型
func (v *Value) Float64() (float64, error) {
	if ret, ok := v.value.(float64); ok {
		return ret, nil
	}
	return strconv.ParseFloat(v.String(), 64)
}

// Int interface{} 转成 int 类型
func (v *Value) Int() (int, error) {
	if ret, ok := v.value.(int); ok {
		return ret, nil
	}
	value, err := strconv.ParseInt(v.String(), 10, 32)
	return int(value), err
}

// Int8 interface{} 转成 int8 类型
func (v *Value) Int8() (int8, error) {
	if ret, ok := v.value.(int8); ok {
		return ret, nil
	}
	value, err := strconv.ParseInt(v.String(), 10, 8)
	return int8(value), err
}

// Int16 interface{} 转成 int16 类型
func (v *Value) Int16() (int16, error) {
	if ret, ok := v.value.(int16); ok {
		return ret, nil
	}
	value, err := strconv.ParseInt(v.String(), 10, 16)
	return int16(value), err
}

// Int32 interface{} 转成 int32 类型
func (v *Value) Int32() (int32, error) {
	if ret, ok := v.value.(int32); ok {
		return ret, nil
	}
	value, err := strconv.ParseInt(v.String(), 10, 32)
	return int32(value), err
}

// Int64 interface{} 转成 int64 类型
func (v *Value) Int64() (int64, error) {
	if ret, ok := v.value.(int64); ok {
		return ret, nil
	}
	value, err := strconv.ParseInt(v.String(), 10, 64)
	return int64(value), err
}

// Uint interface{} 转成 uint 类型
func (v *Value) Uint() (uint, error) {
	if ret, ok := v.value.(uint); ok {
		return ret, nil
	}
	value, err := strconv.ParseUint(v.String(), 10, 32)
	return uint(value), err
}

// Uint8 interface{} 转成 uint8 类型
func (v *Value) Uint8() (uint8, error) {
	if ret, ok := v.value.(uint8); ok {
		return ret, nil
	}
	value, err := strconv.ParseUint(v.String(), 10, 8)
	return uint8(value), err
}

// Uint16 interface{} 转成 uint16 类型
func (v *Value) Uint16() (uint16, error) {
	if ret, ok := v.value.(uint16); ok {
		return ret, nil
	}
	value, err := strconv.ParseUint(v.String(), 10, 16)
	return uint16(value), err
}

// Uint32 interface{} 转成 uint32 类型
func (v *Value) Uint32() (uint32, error) {
	if ret, ok := v.value.(uint32); ok {
		return ret, nil
	}
	value, err := strconv.ParseUint(v.String(), 10, 32)
	return uint32(value), err
}

// Uint64 interface{} 转成 uint64 类型
func (v *Value) Uint64() (uint64, error) {
	if ret, ok := v.value.(uint64); ok {
		return ret, nil
	}
	value, err := strconv.ParseUint(v.String(), 10, 64)
	return uint64(value), err
}

// String interface{} 转成 string 类型
func (v *Value) String() string {
	if ret, ok := v.value.(string); ok {
		return ret
	}
	if v.IsNil() {
		return ""
	}
	return toStr(v.value)
}

// MustBool interface{} 转成 bool 类型
func (v *Value) MustBool() (ret bool) {
	ret, _ = v.Bool()
	return
}

// MustFloat32 convert to float32
func (v *Value) MustFloat32() (ret float32) {
	ret, _ = v.Float32()
	return
}

// MustFloat64 convert to float64
func (v *Value) MustFloat64() (ret float64) {
	ret, _ = v.Float64()
	return
}

// MustInt convert to int
func (v *Value) MustInt() (ret int) {
	ret, _ = v.Int()
	return
}

//MustInt8 convert to int8
func (v *Value) MustInt8() (ret int8) {
	ret, _ = v.Int8()
	return
}

// MustInt16 convert to int16
func (v *Value) MustInt16() (ret int16) {
	ret, _ = v.Int16()
	return
}

// MustInt32 convert to int32
func (v *Value) MustInt32() (ret int32) {
	ret, _ = v.Int32()
	return
}

// MustInt64 convert to int64
func (v *Value) MustInt64() (ret int64) {
	ret, _ = v.Int64()
	return
}

// MustUint convert to uint
func (v *Value) MustUint() (ret uint) {
	ret, _ = v.Uint()
	return
}

// MustUint8 convert to uint8
func (v *Value) MustUint8() (ret uint8) {
	ret, _ = v.Uint8()
	return
}

// MustUint16 convert to uint16
func (v *Value) MustUint16() (ret uint16) {
	ret, _ = v.Uint16()
	return
}

// MustUint32 convert to uint32
func (v *Value) MustUint32() (ret uint32) {
	ret, _ = v.Uint32()
	return
}

// MustUint64 convert to uint64
func (v *Value) MustUint64() (ret uint64) {
	ret, _ = v.Uint64()
	return
}

// MustString convert to string
func (v *Value) MustString() (ret string) {
	ret = v.String()
	return
}

// MustSliceString convert to []string
func (v *Value) MustSliceString() (ret []string) {
	s := reflect.ValueOf(v.value)
	if s.Kind() != reflect.Slice {
		return
	}
	ret = make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = toStr(s.Index(i).Interface())
	}
	return
}

// Must convert to []byte
func (v *Value) MustBytes() (ret []byte) {
	ret, _ = v.value.([]byte)
	return
}

// Value 获取原始的 value
func (v *Value) Value() interface{} {
	return v.value
}

// IsNil 判断 value 是否为 nil
func (v *Value) IsNil() bool {
	return v.value == nil
}

// ToMap 传入接收 map
func (v *Value) ToMap(m interface{}) {

}

// ValueTo 把 interface{} 类型转换成 *Value{} 类型
func ValueTo(value interface{}) *Value {
	return &Value{value: value}
}

// convert interface{} to string
func toStr(value interface{}) (v string) {
	if value == nil {
		return ""
	}
	f := reflect.ValueOf(value)
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = strconv.FormatInt(f.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = strconv.FormatUint(f.Uint(), 10)
	case reflect.Float32:
		v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
	case reflect.Float64:
		v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
	case reflect.String:
		v = f.String()
	case reflect.Bool:
		v = strconv.FormatBool(f.Bool())
	}

	return
}
