package arrays

// 4. Median of Two Sorted Arrays

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := []int{}
	var ai, bj int
	// merge numbers
	for ai < len(nums1) && bj < len(nums2) {
		if nums1[ai] <= nums2[bj] {
			merged = append(merged, nums1[ai])
			ai++
		} else {
			merged = append(merged, nums2[bj])
			bj++
		}
	}
	if ai == len(nums1) {
		merged = append(merged, nums2[bj:]...)
	} else if bj == len(nums2) {
		merged = append(merged, nums1[ai:]...)
	}
	// Calculate merged median
	N := len(merged)
	if N%2 == 0 {
		return float64((float64(merged[int(N/2)-1] + merged[int(N/2)])) / 2)
	}
	return float64(merged[int(N/2)])
}
