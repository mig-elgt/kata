package grind75

import "math"

// Container With Most Water

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Example 1
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

func calculateMaxArea(heights []int) int {
	area := math.MinInt
	lo, hi := 0, len(heights)-1
	for lo < hi {
		length := hi - lo
		height := int(math.Min(float64(heights[hi]), float64(heights[lo])))
		newArea := length * height
		if area < newArea {
			area = newArea
		}
		if heights[hi] <= heights[lo] {
			hi--
		} else {
			lo++
		}
	}
	return area
}
