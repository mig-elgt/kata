package grind75

// Subsets
// Given an integer array nums of unique elements, return all possible subsets (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// Example 2:

// Input: nums = [0]
// Output: [[],[0]]

func getSubsets(nums []int) [][]int {
	subsets := [][]int{{}}
	for _, num := range nums {
		subsetSize := len(subsets)
		for i := 0; i < subsetSize; i++ {
			newSub := make([]int, len(subsets[i]))
			copy(newSub, subsets[i])
			subsets = append(subsets, append(newSub, num))
		}
	}
	return subsets
}
