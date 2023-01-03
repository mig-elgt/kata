package grind75

// Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
	var prefix rune
	foundCommonPrefix := true
	i := 0
	for {
		if i < len(strs[0]) {
			prefix = rune(strs[0][i])
		} else {
			return strs[0][:i]
		}
		for _, str := range strs {
			if i > len(str)-1 {
				return str[0:i]
			}
			if prefix != rune(str[i]) {
				foundCommonPrefix = false
				break
			}
		}
		if !foundCommonPrefix {
			if i > 0 {
				return strs[0][0:i]
			} else {
				return ""
			}
		} else {
			i++
		}
	}
}
