package grind75

// Problem: Maximum Subarray
// Dynamic Programming, Kadane's Algorithm

// Given an integer array nums, find the
// subarray
//  which has the largest sum and return its sum.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Example 2:

// Input: nums = [1]
// Output: 1

import "math"

func maxSubArray(nums []int) int {
	currentSubarray := nums[0]
	maxSubarray := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		currentSubarray = int(math.Max(float64(num), float64(currentSubarray+num)))
		maxSubarray = int(math.Max(float64(maxSubarray), float64(currentSubarray)))
	}
	return maxSubarray
}
