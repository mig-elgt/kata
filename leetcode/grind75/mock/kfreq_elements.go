package mock

import (
	"container/heap"
)

/*
Example 1:
Input : nums = [1,1,1,2,2,3,4,5,6], k = 2
Output: [1,2]

First Aproach
Step 1: Nums Frequency O(n)

Step 2: Sort Map O(nlogn)

	1 -> 3
	2 -> 2
	3 -> 1
	4 -> 1
	5 -> 1
	6 -> 1

Step 3: Loop from i = 0. k O(k)

time complexity = O(n) + O(nlogn) + O(k)

Second Aproach: Min Heap

	2 -> 2
	4 -> 1
	3 -> 1
	6 -> 1
	5 -> 1
	1 -> 3
*/

type NumFreq struct {
	Num, Freq int
}

type MinHeap []NumFreq

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(NumFreq))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	old = old[:s-1]
	*h = old
	return x
}

// Functional Requirements
// Non Functional Requirements
// Api End Points
// Schema Design
// Estimation
// Db selection
// Architecture Diagram(Simple for 1 user)
// --Read Write QPS

func TopKFreqNums(nums []int, k int) []int {
	frequencies := map[int]int{}
	// O(n)
	for _, num := range nums {
		frequencies[num]++
	}
	mh := &MinHeap{}
	heap.Init(mh)
	// O(n)
	for num, freq := range frequencies {
		heap.Push(mh, NumFreq{Num: num, Freq: freq})
		if mh.Len() > k {
			heap.Pop(mh)
		}
	}
	var result []int
	// O(k)
	for mh.Len() > 0 {
		nf := heap.Pop(mh).(NumFreq)
		result = append(result, nf.Num)
	}
	return result
}
