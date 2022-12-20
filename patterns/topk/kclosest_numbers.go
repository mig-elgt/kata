package topk

import (
	"math"
)

func FindClosestElements(nums []int, key, value int) []int {
	panic("not impl")
	// result := []int{}
	// maxHeap := heap.NewHeap(func(a, b int) bool {
	// 	return computeDistance(a, value) < computeDistance(b, value)
	// })
	// for _, num := range nums {
	// 	maxHeap.Push(num)
	// 	if maxHeap.Size() > key {
	// 		maxHeap.Pop()
	// 	}
	// }
	// for maxHeap.Size() > 0 {
	// 	num, _ := maxHeap.Pop()
	// 	result = append(result, num)
	// }
	// return result
}

func computeDistance(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
