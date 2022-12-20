package grind75

// Evaluate Reverse Polish Notation

// Evaluate the value of an arithmetic expression in Reverse Polish Notation.

// Valid operators are +, -, *, and /. Each operand may be an integer or another expression.

// Note that division between two integers should truncate toward zero.

// It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.

// Example 1:

// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

import (
	"container/list"
	"strconv"
)

var ops = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			item := stack.Front()
			b := stack.Remove(item).(int)
			item = stack.Front()
			a := stack.Remove(item).(int)
			res := ops[token](a, b)
			stack.PushFront(res)
		} else {
			num, _ := strconv.Atoi(token)
			stack.PushFront(num)
		}
	}
	item := stack.Front()
	res, _ := stack.Remove(item).(int)
	return res
}
