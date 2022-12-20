package two_pointers

import "math"

func MinimumWindowSort(arr []int) int {
	low, high := 0, len(arr)-1
	for low < len(arr)-1 && arr[low] <= arr[low+1] {
		low++
	}
	if low == len(arr)-1 {
		return 0
	}
	for high > 0 && arr[high] >= arr[high-1] {
		high--
	}
	subMax, subMin := math.MinInt, math.MaxInt
	for k := low; k <= high; k++ {
		subMax = int(math.Max(float64(subMax), float64(arr[k])))
		subMin = int(math.Min(float64(subMin), float64(arr[k])))
	}
	for low > 0 && arr[low-1] > subMin {
		low--
	}
	for high < len(arr)-1 && arr[high+1] < subMax {
		high++
	}
	return high - low + 1
}
