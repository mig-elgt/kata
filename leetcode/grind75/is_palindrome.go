package grind75

import (
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	left, rigth := 0, len(str)-1
	for left <= rigth {
		if !isAlphanumeric(rune(str[left])) {
			left++
			continue
		}
		if !isAlphanumeric(rune(str[rigth])) {
			rigth--
			continue
		}
		if str[left] != str[rigth] {
			return false
		}
		left++
		rigth--
	}
	return true
}

func isAlphanumeric(char rune) bool {
	return (char >= 97 && char <= 122) || (char >= 65 && char <= 90) || (char >= 48 && char <= 57)
}
