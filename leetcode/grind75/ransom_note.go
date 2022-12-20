package grind75

// Ransom Note
// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

// Example 1:

// Input: ransomNote = "a", magazine = "b"
// Output: false

// Example 2:

// Input: ransomNote = "aa", magazine = "aab"
// Output: true

func canConstruct(ransomNote string, magazine string) bool {
	lettersFreq := map[rune]int{}
	for _, letter := range magazine {
		lettersFreq[letter]++
	}
	for _, chr := range ransomNote {
		lettersFreq[chr]--
		if lettersFreq[chr] < 0 {
			return false
		}
	}
	return true
}
