package topk

import "gitlab.com/migel/kata/heap"

func FindKthSmallestNumber(nums []int, k int) int {
	heap := heap.NewHeap(func(a, b int) bool { return a < b })
	i := 0
	for ; i < k; i++ {
		heap.Push(nums[i])
	}
	for ; i < len(nums); i++ {
		num, _ := heap.Peek()
		if num > nums[i] {
			heap.Pop()
			heap.Push(nums[i])
		}
	}
	kth, _ := heap.Peek()
	return kth
}
