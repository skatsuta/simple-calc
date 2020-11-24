package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s '<EXPRESSION>'\n", os.Args[0])
		os.Exit(1)
	}

	// Get a mathematical expression from the first argument
	input := flag.Arg(0)

	// Calculate the expression
	output, err := calc(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print the calculation result to standard output
	if output != nil {
		fmt.Println(output)
	}
}

func calc(input string) (output interface{}, err error) {
	// Your code here...
	return nil, nil
}
