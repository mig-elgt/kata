package slidingwindow

import (
	"container/heap"
	"container/list"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	result := []int{}
	maxHeap := IntHeap{}
	heap.Init(&maxHeap)
	for windowStart, windowEnd := 0, 0; windowEnd < len(nums); windowEnd++ {
		heap.Push(&maxHeap, nums[windowEnd])
		if maxHeap.Len() == k {
			result = append(result, maxHeap[0])
			queue := list.New()
			for maxHeap.Len() > 0 {
				num := heap.Pop(&maxHeap).(int)
				if num != nums[windowStart] {
					queue.PushBack(num)
				} else {
					break
				}
			}
			for queue.Len() > 0 {
				element := queue.Front()
				heap.Push(&maxHeap, element.Value.(int))
				queue.Remove(element)
			}
			windowStart++
		}
	}
	return result
}
