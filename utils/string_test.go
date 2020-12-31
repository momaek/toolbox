package utils

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

func TestRemoveDuplicateItemStringSlice(t *testing.T) {
	s := []string{}
	s1 := RemoveDuplicateStringSlice(s)

	if len(s) != len(s1) {
		t.Fatalf("Expected same length but got len(s) = %d, len(s1) = %d", len(s), len(s1))
	}

	s = []string{"1", "2", "3", "4", "3"}
	s1 = RemoveDuplicateStringSlice(s)
	expectedSlice := []string{"1", "2", "3", "4"}
	for i := range s1 {
		if s1[i] != expectedSlice[i] {
			t.Fatalf("Expect %s but got %s", expectedSlice[i], s1[i])
		}
	}
}
