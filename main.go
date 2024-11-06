package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rahji/wordfinds/internal/wordlist"
	"github.com/spf13/pflag"
)

func main() {

	// define flags
	var (
		inputFile string
		help      bool
	)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
	pflag.BoolVarP(&help, "help", "h", false, "show help message")
	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

	words, err := wordlist.GetWords(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// xxx temporary
	fmt.Println(strings.Join(words, "\n"))
}
