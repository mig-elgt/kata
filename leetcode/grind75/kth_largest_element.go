package grind75

// Problem: Kth Largest Element in an Array

// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// You must solve it in O(n) time complexity.

// Example:

// Input: nums = [3,2,1,5,6,4], k = 2
// Output: 5

import (
	"container/heap"
)

type MinHeapNum []int

func (h MinHeapNum) Len() int           { return len(h) }
func (h MinHeapNum) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapNum) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapNum) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapNum) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minHeap := MinHeapNum{}
	heap.Init(&minHeap)
	for _, num := range nums {
		heap.Push(&minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(&minHeap)
		}
	}
	return heap.Pop(&minHeap).(int)
}
