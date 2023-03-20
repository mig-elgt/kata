package grind75

import (
	"container/heap"
	"fmt"
	"math"
)

type NumberDistance struct {
	Number   int
	Distance int
}

type DistanceMaxHeap []NumberDistance

func (h DistanceMaxHeap) Len() int           { return len(h) }
func (h DistanceMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h DistanceMaxHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance }

func (h *DistanceMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(NumberDistance))
}

func (h *DistanceMaxHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	old = old[:s-1]
	*h = old
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	maxHeap := DistanceMaxHeap{}
	heap.Init(&maxHeap)

	for _, num := range arr {
		dist := int(math.Abs(float64(num - x)))
		fmt.Println(num, dist)
		heap.Push(&maxHeap, NumberDistance{Number: num, Distance: dist})
		if maxHeap.Len() > k {
			heap.Pop(&maxHeap)
		}
	}

	var closests []int
	for maxHeap.Len() > 0 {
		closests = append([]int{heap.Pop(&maxHeap).(NumberDistance).Number}, closests...)
	}
	return closests
}
