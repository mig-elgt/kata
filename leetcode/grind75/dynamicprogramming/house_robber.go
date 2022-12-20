package dynamicprogramming

import "math"

func rob(nums []int) int {
	// return robFrom(0, nums, map[int]int{})
	return robWithTab(nums)
}

// Memoization aproach - Top -> Bottom
func robFrom(i int, nums []int, memo map[int]int) int {
	// No more houses left to examine
	if i >= len(nums) {
		return 0
	}
	if _, cached := memo[i]; cached {
		return memo[i]
	}
	memo[i] = int(math.Max(float64(robFrom(i+1, nums, memo)), float64(robFrom(i+2, nums, memo)+nums[i])))
	return memo[i]
}

// Tabulatin Aproach -> Bottom -> Up
func robWithTab(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	maxRobbedAmount := make([]int, N+1)
	// Base case initialization
	maxRobbedAmount[N] = 0
	maxRobbedAmount[N-1] = nums[N-1]
	// DB Table calculation
	for i := N - 2; i >= 0; i-- {
		maxRobbedAmount[i] = int(math.Max(float64(maxRobbedAmount[i+1]), float64(maxRobbedAmount[i+2]+nums[i])))
	}
	return maxRobbedAmount[0]
}
