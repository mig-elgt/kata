package topk

import "gitlab.com/migel/kata/heap"

func FindTopKFrequentNumbers(nums []int, k int) []int {
	numsFrequency := map[int]int{}
	for _, num := range nums {
		numsFrequency[num]++
	}
	topK := []int{}
	heap := heap.NewHeap(func(a, b int) bool { return numsFrequency[a] > numsFrequency[b] })
	for num := range numsFrequency {
		heap.Push(num)
		if heap.Size() > k {
			heap.Pop()
		}
	}
	for heap.Size() != 0 {
		num, _ := heap.Pop()
		topK = append(topK, num)
	}
	return topK
}
