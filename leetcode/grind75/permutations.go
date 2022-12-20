package grind75

// Permutations

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2:

// Input: nums = [0,1]
// Output: [[0,1],[1,0]]

import (
	"container/list"
)

func permute(nums []int) [][]int {
	results := [][]int{}
	permutations := list.New()
	permutations.PushBack([]int{})
	for _, num := range nums {
		permutationSize := permutations.Len()
		for i := 0; i < permutationSize; i++ {
			item := permutations.Front()
			data := permutations.Remove(item).([]int)
			currentPermutation := make([]int, len(data))
			copy(currentPermutation, data)
			for j := 0; j <= len(currentPermutation); j++ {
				newPermutation := append(currentPermutation[:j], append([]int{num}, currentPermutation[j:]...)...)
				if len(newPermutation) == len(nums) {
					results = append(results, newPermutation)
				}
				permutations.PushBack(newPermutation)
			}
		}
	}
	return results
}
