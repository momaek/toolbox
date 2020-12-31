package utils

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

// RemoveDuplicateStringSlice remove duplicate item in string slice
func RemoveDuplicateStringSlice(slice []string) []string {
	m := make(map[string]struct{})
	retSlice := make([]string, 0, len(slice))
	for _, v := range slice {
		m[v] = struct{}{}
		if len(m) > len(retSlice) {
			retSlice = append(retSlice, v)
		}
	}

	return retSlice
}
