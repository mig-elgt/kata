package grind75

// Problem: Minimum Window Substring
// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring
//  of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

// The testcases will be generated such that the answer is unique.

// Example 1:

// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	uniqueChrFreq := map[rune]int{}
	for _, chr := range t {
		uniqueChrFreq[chr]++
	}
	uniqueChrFreqSize := len(uniqueChrFreq)
	windowCounts := map[rune]int{}
	ans := []int{-1, 0, 0}
	var left, rigth, formed int
	for rigth < len(s) {
		letter := rune(s[rigth])
		windowCounts[letter]++
		if _, ok := uniqueChrFreq[letter]; ok {
			if windowCounts[letter] == uniqueChrFreq[letter] {
				formed++
			}
		}
		for left <= rigth && formed == uniqueChrFreqSize {
			chr := rune(s[left])
			if ans[0] == -1 || rigth-left+1 < ans[0] {
				ans[0] = rigth - left + 1
				ans[1] = left
				ans[2] = rigth
			}
			windowCounts[chr]--
			if _, ok := uniqueChrFreq[chr]; ok {
				if windowCounts[chr] < uniqueChrFreq[chr] {
					formed--
				}
			}
			left++
		}
		rigth++
	}
	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}
