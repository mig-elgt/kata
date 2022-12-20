package sliding_window

import "math"

func SmallestSubArrayWithAGreaterSum(numbers []int, S int) int {
	smallest := math.MaxInt
	var windowSum, windowStart int
	for i := 0; i < len(numbers); i++ {
		windowSum += numbers[i]
		for windowSum >= S {
			if smallest > i-windowStart+1 {
				smallest = i - windowStart + 1
			}
			windowSum -= numbers[windowStart]
			windowStart++
		}

	}
	return smallest
}
