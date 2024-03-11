package lab2

import (
	"fmt"
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePrefix(t *testing.T) {
	testCases := []struct {
		expression string
		expected   int
		err        error
	}{
		{"+ 5 * - 4 2 3", 11, nil},
		{"- * 10 / 20 5 + 3 2", 35, nil},
		{"* + 1 2 + 3 4", 21, nil},
		{"+ 2 2", 4, nil},
		{"- 10 2", 8, nil},
		{"* 10 2", 20, nil},
		{"/ 20 5", 4, nil},
		{"^ 2 3", 8, nil},
		{"", 0, &strconv.NumError{Func: "Atoi", Num: "", Err: strconv.ErrSyntax}},
		{"+ a b", 0, &strconv.NumError{Func: "Atoi", Num: "b", Err: strconv.ErrSyntax}},

	}

	for _, tc := range testCases {
		result, err := CalculatePrefix(tc.expression)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, result)
	}
}

func ExampleCalculatePrefix() {
	res, _ := CalculatePrefix("+ 5 * - 4 2 3")
	fmt.Println(res)

	// Output:
	// 11
}
