package kmerge

import "gitlab.com/migel/kata/heap"

func FindKLargestPairs(l1, l2 []int, k int) [][]int {
	minHeap := heap.NewHeap(func(a, b []int) bool { return a[0]+a[1] > b[0]+b[1] })
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(l2); j++ {
			if minHeap.Size() < k {
				minHeap.Push([]int{l1[i], l2[j]})
			} else {
				sum := l1[i] + l2[j]
				pair, _ := minHeap.Peek()
				if sum < pair[0]+pair[1] {
					break
				} else {
					minHeap.Pop()
					minHeap.Push([]int{l1[i], l2[j]})
				}
			}
		}
	}
	result := [][]int{}
	for minHeap.Size() > 0 {
		pair, _ := minHeap.Pop()
		result = append(result, pair)
	}
	return result
}
