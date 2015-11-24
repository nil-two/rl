package main

type Reverser struct {
}

func NewReverser() *Reverser {
	return &Reverser{}
}

func (r *Reverser) Reverse(s string) string {
	a := []rune(s)
	b := make([]rune, len(a))
	for i := 0; i < len(a); i++ {
		b[len(a)-i-1] = a[i]
	}
	return string(b)
}
