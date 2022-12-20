package two_pointers

import "sort"

// X + Y + Z < target
// Find a pair Y, Z such as Y + Z < target - X
func TripletsWithSmallerSum(arr []int, target int) int {
	sort.Ints(arr)
	var count int
	for i := 0; i < len(arr)-2; i++ {
		targetSum := target - arr[i]
		left, right := i+1, len(arr)-1
		for left < right {
			if arr[left]+arr[right] < targetSum {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}
	return count
}
