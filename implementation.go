package lab2

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func prefixToPostfix(prefixExpression string) (string, error) {

	tokens := strings.Fields(prefixExpression)

	stack := []string{}

	
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for operator")
			}
			first_operand := stack[len(stack)-1]
			second_operand := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			output := fmt.Sprintf("%s %s %s", first_operand, second_operand, token)
			stack = append(stack, output)
		} else if isNumeric(token) {
			stack = append(stack, token)
		} else {
			return "", errors.New("Undefined operator, enter only math operations and numeric values")
		}
	}

	
	if len(stack) != 1 {
		return "", errors.New("error converting expression")
	}

	
	postfixExpression := stack[0]
	return postfixExpression, nil
}