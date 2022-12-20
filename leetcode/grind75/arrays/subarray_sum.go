package arrays

// Pattern: Subarray Sum Equals K

// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,1,1], k = 2
// Output: 2

// Example 2:
// Input: nums = [1,2,3], k = 3
// Output: 2

func subarraySum(nums []int, k int) int {
	freq := map[int]int{0: 1}
	var sum, count int
	for _, num := range nums {
		sum += num
		if _, ok := freq[sum-k]; ok {
			count += freq[sum-k]
		}
		freq[sum] = freq[sum] + 1
	}
	return count
}
