package stringutils

import (
	"reflect"
	"unsafe"
)

// BytesToString ..
func BytesToString(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{Data: bytesHeader.Data, Len: bytesHeader.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func BytesToString2(b []byte) string {
	return string(b)
}
