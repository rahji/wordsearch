package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	// Define flags
	var (
		inputFile string
		help      bool
	)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
	pflag.BoolVarP(&help, "help", "h", false, "show help message")

	// Parse flags
	pflag.Parse()

	// Show help if requested
	if help {
		pflag.Usage()
		os.Exit(0)
	}

	var reader io.Reader

	// If file flag is provided, try to open that file
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		// No file specified, use standard input
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Fprintln(os.Stderr, "No input file specified and no data piped to STDIN")
			fmt.Fprintln(os.Stderr, "Usage: either pipe data to STDIN or use -f flag to specify input file")
			pflag.Usage()
			os.Exit(1)
		}
		reader = os.Stdin
	}

	// Create a scanner to read the input
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	// Process each word
	i := 0
	for scanner.Scan() {
		word := scanner.Text()
		// Process the word here - for now just print it
		i++
		fmt.Printf("%3d %s\n", i, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
