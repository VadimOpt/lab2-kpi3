package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// Create a new scanner with ch.Input as the input
	scanner := bufio.NewScanner(ch.Input)

	// Read the expression from ch.Input
	if scanner.Scan() {
		expression := scanner.Text()

		// Compute the result using CalculatePrefix function
		result, err := CalculatePrefix(expression)
		if err != nil {
			return fmt.Errorf("error computing result: %w", err)
		}

		// Write the result to ch.Output
		_, err = fmt.Fprintln(ch.Output, result)
		if err != nil {
			return fmt.Errorf("error writing result: %w", err)
		}
	} else {
		return fmt.Errorf("error reading expression: %w", scanner.Err())
	}

	return nil
}
