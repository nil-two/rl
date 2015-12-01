package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/yuya-takeyama/argf"
)

func usage() {
	os.Stderr.WriteString(`
Usage: rl [OPTION]... [FILE]...
Reverse lines of FILE(s), or standard input.

Options:
  -d, --delimiter=DELIM    delimit line by DELIM
      --help               display this help text and exit
      --version            output version information and exit
`[1:])
}

func version() {
	os.Stderr.WriteString(`
0.1.0
`[1:])
}

type Option struct {
	Delimiter string `short:"d" long:"delimiter"`
	isHelp    bool   `          long:"help"`
	isVersion bool   `          long:"version"`
	Files     []string
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, "rl:", err)
}

func do(rev *Reverser, r io.Reader) error {
	b := bufio.NewScanner(r)
	for b.Scan() {
		fmt.Println(rev.Reverse(b.Text()))
	}
	return b.Err()
}

func _main() int {
	r, err := argf.From(os.Args[1:])
	if err != nil {
		printErr(err)
		return 2
	}

	rev := NewReverser()
	if err = do(rev, r); err != nil {
		printErr(err)
		return 1
	}
	return 0
}

func main() {
	e := _main()
	os.Exit(e)
}
