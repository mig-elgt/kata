package grind75

// Problem: Word Break

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:
// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".

// Example 2:
// Input: s = "applepenapple", wordDict = ["apple","pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
// Note that you are allowed to reuse a dictionary word.

// Example 3:
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: false

import "container/list"

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := map[string]int{}
	for _, word := range wordDict {
		wordDictSet[word] = 1
	}
	queue := list.New()
	visited := make([]bool, len(s))
	queue.PushBack(0)
	for queue.Len() > 0 {
		element := queue.Front()
		start := element.Value.(int)
		queue.Remove(element)
		if visited[start] {
			continue
		}
		for end := start + 1; end <= len(s); end++ {
			if _, found := wordDictSet[string(s[start:end])]; found {
				queue.PushBack(end)
				if end == len(s) {
					return true
				}
			}
		}
		visited[start] = true
	}
	return false
}
