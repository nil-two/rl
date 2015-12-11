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
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	return strings.Join(a, r.delimiter)
}
