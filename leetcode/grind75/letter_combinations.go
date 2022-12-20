package grind75

// Letter Combinations of a Phone Number

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = ""
// Output: []

// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	letters := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	digitLetters := []string{}
	for _, digit := range digits {
		digitLetters = append(digitLetters, letters[digit])
	}
	result := Result{
		Combinations: []string{},
	}
	backtrackCombinations(0, []byte{}, digits, letters, &result)
	return result.Combinations
}

type Result struct {
	Combinations []string
}

func backtrackCombinations(currentIndex int, path []byte, digits string, digitalLetters map[rune]string, result *Result) {
	if len(path) == len(digits) {
		result.Combinations = append(result.Combinations, string(path))
		return
	}
	letters := digitalLetters[rune(digits[currentIndex])]
	for _, letter := range letters {
		path = append(path, byte(letter))
		backtrackCombinations(currentIndex+1, path, digits, digitalLetters, result)
		path = path[0 : len(path)-1]
	}
}

func backtrackComb(currentComb string, currentDigit int, digitLetters []string, result *Result) {
	if currentDigit == len(digitLetters)-1 {
		for i := 0; i < len(digitLetters[currentDigit]); i++ {
			result.Combinations = append(result.Combinations, currentComb+string(digitLetters[currentDigit][i]))
		}
		return
	}
	for i := 0; i < len(digitLetters[currentDigit]); i++ {
		backtrackComb(currentComb+string(digitLetters[currentDigit][i]), currentDigit+1, digitLetters, result)
	}
}
