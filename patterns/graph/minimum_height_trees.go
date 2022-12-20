package graph

import (
	"gitlab.com/migel/kata/heap"
	"gitlab.com/migel/kata/queue"
)

func FindMimmumHeightTreeNodes(vertices int, edges [][]int) []int {
	graph := map[int][]int{}
	// Init Graph
	for i := 0; i < vertices; i++ {
		graph[i] = []int{}
	}
	// Build Graph given pair edges
	for _, edge := range edges {
		parent, child := edge[0], edge[1]
		children := graph[parent]
		children = append(children, child)
		graph[parent] = children
		children = graph[child]
		children = append(children, parent)
		graph[child] = children
	}
	type node struct {
		vertice int
		heigth  int
	}
	minHeap := heap.NewHeap(func(a, b node) bool { return a.heigth > b.heigth })
	queue := queue.NewQueue(1)
	for i := 0; i < vertices; i++ {
		queue.Push(i)
		visited := make([]bool, vertices)
		var heigth int
		for queue.Len() > 0 {
			heigth++
			size := queue.Len()
			for i := 0; i < size; i++ {
				vertice := queue.Pop()
				visited[vertice] = true
				children := graph[vertice]
				for _, child := range children {
					if !visited[child] {
						queue.Push(child)
					}
				}
			}
		}
		minHeap.Push(node{i, heigth})
	}
	mht := []int{}
	root, _ := minHeap.Pop()
	mht = append(mht, root.vertice)
	for minHeap.Size() > 0 {
		node, _ := minHeap.Pop()
		if node.heigth != root.heigth {
			break
		}
		mht = append(mht, node.vertice)
	}
	return mht
}
