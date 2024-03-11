package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"io"
	lab2 "github.com/forestgreen18/continuous-integration-and-test-automation"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File containing the expression")
	outputFile = flag.String("o", "", "File to write the result")
)

func main() {
	flag.Parse()

	var input io.Reader
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		f, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	} else {
		fmt.Fprintln(os.Stderr, "Please provide an expression or a file containing the expression.")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer f.Close()
		output = f
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error computing result:", err)
		os.Exit(1)
	}
}
