package twopointers

// Problem: Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

func moveZeroes(nums []int) {
	if len(nums) > 1 {
		left, rigth := 0, 1
		for ; rigth < len(nums); rigth++ {
			if nums[left] != 0 {
				left++
			} else if nums[rigth] != 0 {
				nums[left], nums[rigth] = nums[rigth], nums[left]
				left++
			}
		}
	}
}
