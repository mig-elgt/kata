package binarysearch

func OrderAgnosticBinarySearch(arr []int, key int) int {
	start, end := 0, len(arr)-1
	isAscending := arr[start] < arr[end]
	for start <= end {
		middle := start + (end-start)/2
		if key == arr[middle] {
			return middle
		}
		if key > arr[middle] {
			if isAscending {
				start = middle + 1
			} else {
				end = middle - 1
			}
		} else {
			if isAscending {
				end = middle - 1
			} else {
				start = middle + 1
			}
		}
	}
	return -1
}
