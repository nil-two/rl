package main

import (
	"strings"
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

var reverseWithDelimiterTests = []struct {
	delimiter string
	src       string
	dst       string
}{
	{",", "abcd", "abcd"},
	{",", "10,200,3000", "3000,200,10"},
	{"、", "十、二百、三千", "三千、二百、十"},
}

func TestReverseWithDelimiter(t *testing.T) {
	rev := NewReverser()
	for _, test := range reverseWithDelimiterTests {
		rev.SetDelimiter(test.delimiter)
		expect := test.dst
		actual := rev.Reverse(test.src)
		if actual != expect {
			t.Errorf("%q.Reverse(%q) = %q, want %q",
				rev.delimiter, test.src, actual, expect)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	src := strings.Repeat("abc", 10000)
	rev := NewReverser()
	for i := 0; i < b.N; i++ {
		rev.Reverse(src)
	}
}

func BenchmarkReverseWithDelimiter(b *testing.B) {
	src := strings.TrimSuffix(strings.Repeat("abc,", 10000), ",")
	rev := NewReverser()
	rev.SetDelimiter(",")
	for i := 0; i < b.N; i++ {
		rev.Reverse(src)
	}
}
