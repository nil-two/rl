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
	name    = "rl"
	version = "0.3.1"

	flagset   = pflag.NewFlagSet(name, pflag.ContinueOnError)
	delimiter = flagset.StringP("delimiter", "d", "", "")
	isHelp    = flagset.BoolP("help", "h", false, "")
	isVersion = flagset.BoolP("version", "v", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]... [FILE]...
Reverse lines of FILE(s), or standard input.

Options:
  -d, --delimiter=DELIM    delimit lines by DELIM
  -h, --help               display this help text and exit
  -v, --version            output version information and exit
`[1:], name)
}

func printVersion() {
	fmt.Fprintln(os.Stderr, version)
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
}

func guideToHelp() {
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", name)
}

func do(rev *Reverser, r io.Reader) error {
	b := bufio.NewScanner(r)
	for b.Scan() {
		fmt.Println(rev.Reverse(b.Text()))
	}
	return b.Err()
}

func _main() int {
	flagset.SetOutput(ioutil.Discard)
	if err := flagset.Parse(os.Args[1:]); err != nil {
		printErr(err)
		guideToHelp()
		return 2
	}
	switch {
	case *isHelp:
		printUsage()
		return 0
	case *isVersion:
		printVersion()
		return 0
	}

	r, err := argf.From(flagset.Args())
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
