package utils

// Pointer ..
func Pointer[T any](val T) *T {
	return &val
}
