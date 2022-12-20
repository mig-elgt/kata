package grind75

import (
	"math"
	"sort"
)

// Merge Intervals

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if len(merged) == 0 || merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
		} else {
			a := merged[len(merged)-1]
			b := intervals[i]
			merged[len(merged)-1][0] = int(math.Min(float64(a[0]), float64(b[0])))
			merged[len(merged)-1][1] = int(math.Max(float64(a[1]), float64(b[1])))
		}
	}
	return merged
}
