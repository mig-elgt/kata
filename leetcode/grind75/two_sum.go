package grind75

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example 1
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func twoSum(nums []int, target int) []int {
	indices := map[int]int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if value, ok := indices[complement]; ok {
			return []int{value, i}
		}
		indices[nums[i]] = i
	}
	return []int{-1, -1}
}
