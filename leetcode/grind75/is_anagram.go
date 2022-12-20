package grind75

// Valid Anagram

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false

func isAnagram(s, t string) bool {
	frequency := map[rune]int{}
	for _, chr := range s {
		frequency[chr]++
	}
	for _, chr := range t {
		frequency[chr]--
		if frequency[chr] == 0 {
			delete(frequency, chr)
		}
	}
	if len(frequency) > 0 {
		return false
	}
	return true
}
