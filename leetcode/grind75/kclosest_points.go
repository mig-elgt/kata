package grind75

// Problem: K Closest Points to Origin

// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).
// The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
// You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).

// Example

// Input: points = [[1,3],[-2,2]], k = 1
// Output: [[-2,2]]
// Explanation:
// The distance between (1, 3) and the origin is sqrt(10).
// The distance between (-2, 2) and the origin is sqrt(8).
// Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
// We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].

import "container/heap"

type PointDistance struct {
	index int
	value int
}

type MaxHeap []PointDistance

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(PointDistance))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for i := 0; i < len(points); i++ {
		heap.Push(maxHeap, PointDistance{
			index: i, value: eucludean(points[i]),
		})
		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}
	closestPoints := [][]int{}
	for maxHeap.Len() > 0 {
		index := heap.Pop(maxHeap).(PointDistance).index
		closestPoints = append(closestPoints, points[index])
	}
	return closestPoints
}

func eucludean(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
