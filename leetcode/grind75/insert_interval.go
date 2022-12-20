package grind75

import (
	"math"
	"sort"
)

// Insert Intervals

// You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

// Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

// Return intervals after the insertion.

// Example 1:

// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

func insertInteval(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	for i := 0; i < len(intervals); {
		a := intervals[i]
		j := i + 1
		for ; j < len(intervals); j++ {
			b := intervals[j]
			if b[0] <= a[1] {
				a[0] = int(math.Min(float64(a[0]), float64(b[0])))
				a[1] = int(math.Max(float64(a[1]), float64(b[1])))
			} else {
				break
			}
		}
		result = append(result, a)
		i = j
	}
	return result
}
