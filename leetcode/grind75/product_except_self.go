package grind75

// Product of Array Except Self

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

func calculateProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i <= len(nums)-1; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	suffix := 1
	for i := len(nums) - 1; i > 0; i-- {
		answer[i-1] *= suffix * nums[i]
		suffix *= nums[i]
	}
	return answer
}
