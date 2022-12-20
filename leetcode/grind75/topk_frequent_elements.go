package grind75

// Top K Frequent Elements
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

import "container/heap"

type NumFrequent struct {
	num, frequency int
}

type MinHeap []NumFrequent

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(NumFrequent))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentElements(nums []int, k int) []int {
	numsFrequency := map[int]int{}
	for _, n := range nums {
		numsFrequency[n]++
	}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for num, freq := range numsFrequency {
		heap.Push(minHeap, NumFrequent{
			num:       num,
			frequency: freq,
		})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	result := []int{}
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(minHeap).(NumFrequent).num)
	}
	return result
}
