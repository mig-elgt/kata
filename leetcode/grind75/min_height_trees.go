package grind75

// 310. Minimum Height Trees

// A tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.

// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).

// Return a list of all MHTs' root labels. You can return the answer in any order.

// The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.

// Example 1
// Input: n = 4, edges = [[1,0],[1,2],[1,3]]
// Output: [1]
// Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.

// Example 2
// Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
// Output: [3,4]
// Constraints:

// 1 <= n <= 2 * 104
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// All the pairs (ai, bi) are distinct.
// The given input is guaranteed to be a tree and there will be no repeated edges.

import (
	"container/heap"
	"container/list"
	"sync"
)

type NodeHeight struct {
	Node   int
	Height int
}

type minHeightHeap []NodeHeight

func (h minHeightHeap) Less(i, j int) bool { return h[i].Height < h[j].Height }
func (h minHeightHeap) Len() int           { return len(h) }
func (h minHeightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

var mu sync.Mutex

func (h *minHeightHeap) Push(x interface{}) {
	mu.Lock()
	*h = append(*h, x.(NodeHeight))
	mu.Unlock()
}

func (h *minHeightHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	old = old[:s-1]
	*h = old
	return x
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 || len(edges) == 0 {
		return []int{0}
	}
	tree := map[int][]int{}
	for i := 0; i < n; i++ {
		tree[i] = []int{}
	}
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}
	minHeap := minHeightHeap{}
	heap.Init(&minHeap)
	var wg sync.WaitGroup
	wg.Add(n)
	for root := 0; root < n; root++ {
		root := root
		go func() {
			defer wg.Done()
			visited := map[int]int{}
			height := 0
			queue := list.New()
			queue.PushBack(root)
			for queue.Len() > 0 {
				size := queue.Len()
				for i := 0; i < size; i++ {
					item := queue.Front()
					node := item.Value.(int)
					for _, n := range tree[node] {
						if _, ok := visited[n]; !ok {
							queue.PushBack(n)
						}
					}
					queue.Remove(item)
					visited[node] = 1
				}
				height++
			}
			heap.Push(&minHeap, NodeHeight{Node: root, Height: height})
		}()
	}
	wg.Wait()
	results := []int{}
	for minHeap.Len() > 0 {
		results = append(results, minHeap[0].Node)
		node := heap.Pop(&minHeap).(NodeHeight)
		if minHeap.Len() > 0 && node.Height != minHeap[0].Height {
			break
		}
	}
	return results
}
