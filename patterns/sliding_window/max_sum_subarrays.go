package sliding_window

import "math"

func FindMaxSumSubArrays(numbers []int, k int) int {
	maxSum := math.MinInt
	var sum, j int
	for i := 0; i < len(numbers)-1; i++ {
		sum += numbers[i]
		if i >= k-1 {
			if maxSum < sum {
				maxSum = sum
			}
			sum -= numbers[j]
			j++
		}
	}
	return maxSum
}
