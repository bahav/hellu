package main

import (
	"testing"
)

func TestSomeFunc(t *testing.T) {
	tests := [][]string{
		[]string{"a", "aa"},
		[]string{"ab", "abab"},
		[]string{"", ""},
	}

	for _, a := range tests {
		got := someFunc(a[0])
		if got != a[1] {
			t.Errorf("Failed for: %s. Got: %s", a[0], got)
		}
	}
}
