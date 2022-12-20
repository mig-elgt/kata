package grind75

// 3Sum

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.
//
// Example 1:

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			triplets = append(triplets, twoSumII(nums, i)...)
		}
	}
	return triplets
}

func twoSumII(nums []int, i int) [][]int {
	res := [][]int{}
	lo, hi := i+1, len(nums)-1
	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		if sum < 0 {
			lo++
		} else if sum > 0 {
			hi--
		} else {
			res = append(res, []int{nums[i], nums[lo], nums[hi]})
			lo++
			hi--
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}
		}
	}
	return res
}
