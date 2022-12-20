package slidingwindow

// Problem:  Longest Repeating Character Replacement
// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:

// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
// Example 2:

// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.

func characterReplacement(s string, k int) int {
	var start, maxFrequency, longestSubStrLen int
	freqMap := make([]int, 26)
	for end := 0; end < len(s); end++ {
		currentChr := s[end] - 'A'
		freqMap[currentChr]++
		if freqMap[currentChr] > maxFrequency {
			maxFrequency = freqMap[currentChr]
		}
		if !(end+1-start-maxFrequency <= k) {
			outgoingChair := s[start] - 'A'
			freqMap[outgoingChair]--
			start++
		}
		longestSubStrLen = end + 1 - start
	}
	return longestSubStrLen
}

func binarySearchAndSlidingWindow(s string, k int) int {
	lo, hi := 1, len(s)+1
	for lo+1 < hi {
		mid := lo + (hi-lo)/2
		if isValidSubString(s, mid, k) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func isValidSubString(str string, subStrLen, k int) bool {
	freqMap := make([]int, 26)
	var start, maxFrequency int
	for end := 0; end < len(str); end++ {
		freqMap[str[end]-'A']++
		if end+1-start > subStrLen {
			freqMap[str[end]-'A']--
			start++
		}
		if freqMap[str[end]-'A'] > maxFrequency {
			maxFrequency = freqMap[str[end]-'A']
		}
		if subStrLen-maxFrequency <= k {
			return true
		}
	}
	return false
}
