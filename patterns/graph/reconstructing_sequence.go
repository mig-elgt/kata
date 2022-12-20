package graph

import "gitlab.com/migel/kata/queue"

func CanConstructOriginalSequence(originalSeq []int, sequences [][]int) bool {
	graph := map[int][]int{}
	inDegree := map[int]int{}
	// Init graph and inDegre
	for _, seq := range sequences {
		for _, num := range seq {
			graph[num] = []int{}
			inDegree[num] = 0
		}
	}
	// Build graph and inDegree
	for _, seq := range sequences {
		for i := 0; i < len(seq)-1; i++ {
			parent, child := seq[i], seq[i+1]
			children := graph[parent]
			children = append(children, child)
			graph[parent] = children
			inDegree[child]++
		}
	}
	queue := queue.NewQueue(1)
	// find node sources
	for node, edges := range inDegree {
		if edges == 0 {
			queue.Push(node)
		}
	}
	seqIndex := 0
	for queue.Len() > 0 {
		if queue.Len() > 1 {
			return false
		}
		num := queue.Pop()
		if originalSeq[seqIndex] != num {
			return false
		}
		children := graph[num]
		for _, child := range children {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue.Push(child)
			}
		}
		seqIndex++
	}
	if seqIndex < len(originalSeq) {
		return false
	}
	return true
}
