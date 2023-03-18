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

func main() {
	
	prefixExpression1 := "^ 2 3.5"
	prefixExpression2 := "* + 2 3 4"
	prefixExpression3 := "- / * 3 4 + 5 6 7"
	prefixExpression4 := "+ 1 4 + +"

	
	postfixExpression1, err1 := prefixToPostfix(prefixExpression1)
	postfixExpression2, err2 := prefixToPostfix(prefixExpression2)
	postfixExpression3, err3 := prefixToPostfix(prefixExpression3)
	postfixExpression4, err4 := prefixToPostfix(prefixExpression4)

	
	fmt.Println("Input expression 1:", prefixExpression1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Output expression 1:", postfixExpression1)
	}

	fmt.Println("Input expression 2:", prefixExpression2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Output expression 2:", postfixExpression2)
	}

	fmt.Println("Input expression 3:", prefixExpression3)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Output expression 3:", postfixExpression3)
	}

	fmt.Println("Input expression 4:", prefixExpression4)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Println("Output expression 4:", postfixExpression4)
	}
}