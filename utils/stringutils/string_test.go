package stringutils

import (
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	by := []byte("helloworld")
	for i := 0; i < b.N; i++ {
		BytesToString(by)
	}
}

func BenchmarkBytesToString2(b *testing.B) {
	by := []byte("helloworld")
	for i := 0; i < b.N; i++ {
		BytesToString2(by)
	}
}
