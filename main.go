package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/yuya-takeyama/argf"
)

func printErr(err error) {
	fmt.Fprintln(os.Stderr, "rl:", err)
}

func do(r io.Reader) error {
	b := bufio.NewScanner(r)
	for b.Scan() {
		fmt.Println(reverse(b.Text()))
	}
	return b.Err()
}

func _main() int {
	r, err := argf.From(os.Args[1:])
	if err != nil {
		printErr(err)
		return 2
	}

	if err = do(r); err != nil {
		printErr(err)
		return 1
	}
	return 0
}

func main() {
	e := _main()
	os.Exit(e)
}
