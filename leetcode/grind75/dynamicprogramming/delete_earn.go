package dynamicprogramming

import "math"

func deleteAndEarn(nums []int) int {
	points := map[int]int{}
	maxNum := -1
	for _, num := range nums {
		points[num] += num
		if num > maxNum {
			maxNum = num
		}
	}
	// return maxPoints(maxNum, points, map[int]int{})
	return deleteAndEarnBottomUp(nums, points, maxNum)
}

func deleteAndEarnBottomUp(nums []int, points map[int]int, maxNumber int) int {
	// Base case
	twoBack, oneBack := 0, points[1]
	for num := 2; num <= maxNumber; num++ {
		tmp := oneBack
		oneBack = int(math.Max(float64(oneBack), float64(twoBack+points[num])))
		twoBack = tmp
	}
	return oneBack
}

// top-down approach
func maxPoints(num int, points, cache map[int]int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return points[num]
	}
	if _, found := cache[num]; found {
		return cache[num]
	}
	gain := points[num]
	cache[num] = int(math.Max(float64(maxPoints(num-1, points, cache)), float64(maxPoints(num-2, points, cache)+gain)))
	return cache[num]
}
