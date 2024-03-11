package lab2


import (
	"fmt"
	"strconv"
	"strings"
)

func CalculatePrefix(prefix string) (int, error) {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"^": true,
	}

	stack := []int{}

	// Split the prefix string into tokens
	tokens := strings.Split(prefix, " ")

	// Iterate over the tokens from right to left
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if !operators[token] {
			// If the token is a number, push it onto the stack
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			// If the token is an operator, pop two numbers from the stack, perform the operation, and push the result back onto the stack
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid prefix expression")
			}
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			case "^":
				result := 1
				for i := 0; i < num2; i++ {
					result *= num1
				}
				stack = append(stack, result)
			}
		}
	}

	// If there's exactly one number left on the stack, that's the result. Otherwise, the prefix expression was invalid.
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid prefix expression")
	}
	return stack[0], nil
}
