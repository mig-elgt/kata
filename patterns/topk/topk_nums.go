package topk

import "gitlab.com/migel/kata/heap"

func FindKLargestNumbers(nums []int, k int) []int {
	h := heap.NewHeap(func(a, b int) bool { return a > b })
	largestNums := []int{}
	for i := 0; i < len(nums); i++ {
		if i < 3 {
			h.Push(nums[i])
		} else {
			root, _ := h.Peek()
			if root < nums[i] {
				h.Pop()
				h.Push(nums[i])
			}
		}
	}
	for {
		num, empty := h.Pop()
		if empty {
			largestNums = append(largestNums, num)
		} else {
			break
		}
	}
	return largestNums
}
