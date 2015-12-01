package main

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
	a := []rune(s)
	b := make([]rune, len(a))
	for i := 0; i < len(a); i++ {
		b[len(a)-i-1] = a[i]
	}
	return string(b)
}
