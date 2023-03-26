package grind75

import "container/heap"

type MaxSumHeap []int

func (h MaxSumHeap) Len() int           { return len(h) }
func (h MaxSumHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxSumHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MaxSumHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxSumHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	old = old[:s-1]
	*h = old
	return x
}

func findMaxSum(nums []int) int {
	minHeap := MaxSumHeap{}
	heap.Init(&minHeap)
	for _, num := range nums {
		heap.Push(&minHeap, num)
		if minHeap.Len() > 2 {
			heap.Pop(&minHeap)
		}
	}
	var maxSum int
	for minHeap.Len() > 0 {
		maxSum += heap.Pop(&minHeap).(int)
	}
	return maxSum
}
