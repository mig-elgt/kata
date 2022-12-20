package grind75

import (
	"container/list"
	"math"
)

func largestRectangleArea(heights []int) int {
	stack := list.New()
	stack.PushBack(-1)
	maxArea, n := 0, len(heights)
	for i := 0; i < n; i++ {
		for {
			item := stack.Back()
			itemValue := item.Value.(int)
			if itemValue != -1 && heights[itemValue] >= heights[i] {
				currentHeight := heights[itemValue]
				stack.Remove(item)
				currentWidth := i - stack.Back().Value.(int) - 1
				maxArea = int(math.Max(float64(maxArea), float64(currentHeight*currentWidth)))
			} else {
				break
			}
		}
		stack.PushBack(i)
	}
	for {
		item := stack.Back()
		itemValue := item.Value.(int)
		if itemValue != -1 {
			currentHeight := heights[itemValue]
			stack.Remove(item)
			currentWidth := n - stack.Back().Value.(int) - 1
			maxArea = int(math.Max(float64(maxArea), float64(currentHeight*currentWidth)))
		} else {
			break
		}
	}
	return maxArea
}
