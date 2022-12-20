package twopointers

// Problem: 3Sum Closest

// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers.
// You may assume that each input would have exactly one solution.

// Example 1:
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

// Example 2:
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var minSum int
	minDist := math.MaxInt
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			dist := int(math.Abs(float64(target - sum)))
			if dist == 0 {
				return sum
			}
			if dist < minDist {
				minSum = sum
				minDist = dist
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return minSum
}
