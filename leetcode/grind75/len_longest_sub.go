package grind75

import "math"

// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	letters := map[rune]int{}
	longest := math.MinInt
	windowStart := 0
	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		chr := rune(s[windowEnd])
		if _, ok := letters[chr]; ok {
			for windowStart < windowEnd {
				chrToDelete := rune(s[windowStart])
				letters[chrToDelete]--
				if letters[chrToDelete] == 0 {
					delete(letters, chrToDelete)
				}
				windowStart++
				if chrToDelete == chr {
					letters[chr]++
					break
				}
			}
		} else {
			letters[chr]++
			if longest < len(letters) {
				longest = len(letters)
			}
		}
	}
	return longest
}
