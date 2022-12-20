package arrays

// Problem: Meeting Rooms II

// Given an array of meeting time intervals intervals where intervals[i] = [starti, endi], return the minimum number of conference rooms required.

// Example 1:
// Input: intervals = [[0,30],[5,10],[15,20]]
// Output: 2

// Example 2:
// Input: intervals = [[7,10],[2,4]]
// Output: 1

import (
	"container/heap"
	"sort"
)

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	*h = old[:s-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	minHeap := MinIntHeap{}
	heap.Init(&minHeap)
	heap.Push(&minHeap, intervals[0][1])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= minHeap[0] {
			heap.Pop(&minHeap)
		}
		heap.Push(&minHeap, intervals[i][1])
	}
	return minHeap.Len()
}
