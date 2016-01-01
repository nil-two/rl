package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/ogier/pflag"
	"github.com/yuya-takeyama/argf"
)

var (
	version = "0.2.1"

	flag      = pflag.NewFlagSet("rl", pflag.ContinueOnError)
	delimiter = flag.StringP("delimiter", "d", "", "")
	isHelp    = flag.BoolP("help", "", false, "")
	isVersion = flag.BoolP("version", "", false, "")
)

func usage() {
	os.Stderr.WriteString(`
Usage: rl [OPTION]... [FILE]...
Reverse lines of FILE(s), or standard input.

Options:
  -d, --delimiter=DELIM    delimit lines by DELIM
      --help               display this help text and exit
      --version            output version information and exit
`[1:])
}

func printVersion() {
	fmt.Fprintln(os.Stderr, version)
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, "rl:", err)
}

func guideToHelp() {
	os.Stderr.WriteString(`
Try 'rl --help' for more information.
`[1:])
}

func do(rev *Reverser, r io.Reader) error {
	b := bufio.NewScanner(r)
	for b.Scan() {
		fmt.Println(rev.Reverse(b.Text()))
	}
	return b.Err()
}

func _main() int {
	flag.SetOutput(ioutil.Discard)
	if err := flag.Parse(os.Args[1:]); err != nil {
		printErr(err)
		guideToHelp()
		return 2
	}
	switch {
	case *isHelp:
		usage()
		return 0
	case *isVersion:
		printVersion()
		return 0
	}

	r, err := argf.From(flag.Args())
	if err != nil {
		printErr(err)
		return 2
	}
	rev := NewReverser()
	rev.SetDelimiter(*delimiter)
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
