package mergeintervals

import "math"

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	mergedIntervals := [][]int{}
	i := 0
	// skip (and add to output) all intervals that come before the 'newInterval'
	for i < len(intervals) && intervals[i][0] < newInterval[0] {
		mergedIntervals = append(mergedIntervals, intervals[i])
		i++
	}
	// merge all intervals that overlap with 'newInterval'
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = int(math.Min(float64(intervals[i][0]), float64(newInterval[0])))
		newInterval[1] = int(math.Max(float64(intervals[i][1]), float64(newInterval[1])))
		i++
	}
	mergedIntervals = append(mergedIntervals, newInterval)
	for i < len(intervals) {
		mergedIntervals = append(mergedIntervals, intervals[i])
		i++
	}
	return mergedIntervals
}
