package grind75

// Majority Element
// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func majorityElement(nums []int) int {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
		if frequency[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}
