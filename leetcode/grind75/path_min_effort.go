package grind75

import (
	"container/heap"
	"math"
)

type heightDifference struct {
	X, Y, Diff int
}

type HeightMinHeap []heightDifference

func (h HeightMinHeap) Len() int           { return len(h) }
func (h HeightMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h HeightMinHeap) Less(i, j int) bool { return h[i].Diff > h[j].Diff }

func (h *HeightMinHeap) Push(x interface{}) {
	*h = append(*h, x.(heightDifference))
}

func (h *HeightMinHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	old = old[0 : s-1]
	*h = old
	return x
}

func minimumEffortPath(heights [][]int) int {
	visited := [][]bool{}
	for i := 0; i < len(heights); i++ {
		visited = append(visited, make([]bool, len(heights[0])))
	}

	minHeap := HeightMinHeap{}
	heap.Init(&minHeap)
	heap.Push(&minHeap, &heightDifference{0, 0, 0})

	for minHeap.Len() > 0 {
		hd := heap.Pop(&minHeap).(heightDifference)
		if visited[hd.X][hd.Y] && hd.X == len(heights) && hd.Y == len(heights[0]) {
			return heap.Pop(&minHeap).(heightDifference).Diff
		}
		visited[hd.X][hd.Y] = true
		i, j := hd.X, hd.Y
		if i+1 < len(heights) && j+1 < len(heights[0]) {
			// Move to right
			if visited[i][j+1] {
				diff := int(math.Abs(float64(heights[i][j+1] - hd.Diff)))
				heap.Push(&minHeap, heightDifference{i, j + 1, diff})
			}
			// Move to bottom
			if visited[i+1][j] {
				diff := int(math.Abs(float64(heights[i+1][j] - heights[i][j])))
				heap.Push(&minHeap, heightDifference{i + 1, j, diff})
			}
			// Move to left
			if j-1 >= 0 && visited[i][j-1] {
				diff := int(math.Abs(float64(heights[i][j-1] - heights[i][j])))
				heap.Push(&minHeap, heightDifference{i, j - 1, diff})
			}
			// Move to up
			if i-1 >= 0 && visited[i-1][j] {
				diff := int(math.Abs(float64(heights[i-1][j] - heights[i][j])))
				heap.Push(&minHeap, heightDifference{i - 1, j, diff})
			}
		}

	}
	return 0
}
