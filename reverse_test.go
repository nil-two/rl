package main

import (
	"testing"
)

var reverseDefaultTests = []struct {
	src string
	dst string
}{
	{"abcd", "dcba"},
	{" \t\n", "\n\t "},
	{"日本語", "語本日"},
	{"10,200,3000", "0003,002,01"},
}

func TestReverseDefault(t *testing.T) {
	rev := NewReverser()
	for _, test := range reverseDefaultTests {
		expect := test.dst
		actual := rev.Reverse(test.src)
		if actual != expect {
			t.Errorf("Reverse(%q) = %q, want %q",
				test.src, actual, expect)
		}
	}
}
