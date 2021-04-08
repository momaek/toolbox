package iputils

import (
	"net"
	"testing"
)

func TestFastParseIP(t *testing.T) {
	expected := "1.22.44.255"
	if got := ParseIPv4("1.22.44.255").String(); got != expected {
		t.Fatalf("unexpected response, expected: %s, got: %s", expected, got)
		return
	}

}

func BenchmarkNetParseIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		net.ParseIP("2.2.2.2")
	}
}

func BenchmarkFastParseIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseIPv4("2.2.2.2")
	}
}
