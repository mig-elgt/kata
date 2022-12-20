package binarysearch

func RotationCount(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		middle := start + (end-start)/2
		if arr[start] <= arr[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	if start < len(arr)-1 {
		return start + 1
	}
	return 0
}
