package arrays

// Problem:  Group Anagrams
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for i, str := range strs {
		word := []rune(str)
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		groups[string(word)] = append(groups[string(word)], strs[i])
	}
	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
