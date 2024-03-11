package main

import (
	"flag"
	"fmt"
	lab2 "github.com/forestgreen18/continuous-integration-and-test-automation"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	result, err := lab2.CalculatePrefix("+ 5 * - 4 2 3")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()
		lab2.TestCalculatePrefix();
		lab2.ExampleCalculatePrefix();

	// res, _ := lab2.PrefixToPostfix("+ 2 2")
	// fmt.Println(res)
}
