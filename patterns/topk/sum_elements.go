package topk

import "gitlab.com/migel/kata/heap"

func FindSumOfElements(nums []int, k1, k2 int) int {
	heap := heap.NewHeap(func(a, b int) bool { return a > b })
	for _, num := range nums {
		heap.Push(num)
	}
	for i := 0; i < k1; i++ {
		heap.Pop()
	}
	var sum int
	for i := 0; i < k2-k1-1; i++ {
		num, _ := heap.Pop()
		sum += num
	}
	return sum
}
