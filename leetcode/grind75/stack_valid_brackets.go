package grind75

// Valid Parentheses
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false

func isValid(str string) bool {
	stack := []rune{}
	for _, chr := range str {
		if chr == '(' || chr == '{' || chr == '[' {
			stack = append(stack, chr)
		} else {
			if len(stack) == 0 {
				return false
			}
			tileIndex := len(stack) - 1
			if chr == ')' && stack[tileIndex] != '(' || chr == '}' && stack[tileIndex] != '{' || chr == ']' && stack[tileIndex] != '[' {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
