package grind75

// 394. Decode String
// Given an encoded string, return its decoded string.

// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

// The test cases are generated so that the length of the output will never exceed 105.

// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:
// Input: s = "3[a2[c]]"
// Output: "accaccacc"

// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"

// Constraints:

// 1 <= s.length <= 30
// s consists of lowercase English letters, digits, and square brackets '[]'.
// s is guaranteed to be a valid input.
// All the integers in s are in the range [1, 300].

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []interface{}

func (this *Stack) Push(x interface{}) {
	*this = append(*this, x)
}

func (this *Stack) Pop() interface{} {
	if len(*this) == 0 {
		return nil
	}
	old := *this
	s := len(old)
	x := old[s-1]
	old = old[:s-1]
	*this = old
	return x
}

func decodeString(s string) string {
	str := []rune(s)
	timesStack := Stack{}
	encodedStack := Stack{}
	for i := 0; i < len(str); i++ {
		var token string
		for isNumber(str[i]) {
			token += string(str[i])
			i++
		}
		if token != "" { // found a number
			num, _ := strconv.Atoi(token)
			timesStack.Push(num)
		}
		if str[i] != ']' {
			encodedStack.Push(string(str[i]))
			continue
		}
		times := timesStack.Pop().(int)
		var subDecodeStr string
		for {
			x := encodedStack.Pop()
			if x == nil {
				break
			}
			if x.(string) == "[" {
				break
			}
			subDecodeStr = fmt.Sprintf("%v%v", x, subDecodeStr)
		}
		res := strings.Repeat(subDecodeStr, times)
		encodedStack.Push(res)
	}
	var result string
	for {
		x := encodedStack.Pop()
		if x == nil {
			break
		}
		result = fmt.Sprintf("%v%v", x.(string), result)
	}
	return result
}

func isNumber(chr rune) bool {
	return chr >= '0' && chr <= '9'
}
