package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumeric(t *testing.T) {
    assert.True(t, isNumeric("123"))
    assert.True(t, isNumeric("0.5"))
    assert.True(t, isNumeric("-10"))
    assert.False(t, isNumeric("abc"))
    assert.False(t, isNumeric("123a"))
}

func TestPrefixToPostfix(t *testing.T) {
    testCases := []struct {
        prefixExpression  string
        postfixExpression string
        err         error
    }{
        {"+ 2 3.5", "2 3.5 +", nil},
        {"* + 2 3 4", "2 3 + 4 *", nil},
        {"- / * 3 4 + 5 6 7", "3 4 * 5 6 + / 7 -", nil},
        {"+ 1 4 + +", "", errors.New("insufficient operands for operator")},
        {"+", "", errors.New("insufficient operands for operator")},
        {"a b c", "", errors.New("Undefined operator, enter only math operations and numeric values")},
    }

    for _, tc := range testCases {
        postfixExpression, err := prefixToPostfix(tc.prefixExpression)
        assert.Equal(t, tc.err, err)
        assert.Equal(t, tc.postfixExpression, postfixExpression)
    }

}

func ExamplePrefixToPostfix() {
	res, err := prefixToPostfix("+ 3 4")
	if err == nil {
		fmt.Println(res)
	} else {
		panic(err)
	}
}
