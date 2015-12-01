package main

import (
	"strings"
)

type Reverser struct {
	delimiter string
}

func NewReverser() *Reverser {
	return &Reverser{}
}

func (r *Reverser) SetDelimiter(s string) {
	r.delimiter = s
}

func (r *Reverser) Reverse(s string) string {
	a := strings.Split(s, r.delimiter)
	b := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		b[len(a)-i-1] = a[i]
	}
	return strings.Join(b, r.delimiter)
}
