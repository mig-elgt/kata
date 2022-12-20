package grind75

import "container/list"

// Clone Graph
// Given a reference of a node in a connected undirected graph.
// Return a deep copy (clone) of the graph.

// Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

// class Node {
//     public int val;
//     public List<Node> neighbors;
// }

// Test case format:

// For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

// The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

// Example 1:
// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
// Output: [[2,4],[1,3],[2,4],[1,3]]
// Explanation: There are 4 nodes in the graph.
// 1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
// 3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeRefs := map[*Node]*Node{}
	nodeRefs[node] = &Node{Val: node.Val}
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			e := queue.Front()
			queue.Remove(e)
			currentNode := e.Value.(*Node)
			for _, neighbor := range currentNode.Neighbors {
				if _, ok := nodeRefs[neighbor]; !ok {
					nodeRefs[neighbor] = &Node{Val: neighbor.Val}
					queue.PushBack(neighbor)
				}
				nodeRefs[currentNode].Neighbors = append(nodeRefs[currentNode].Neighbors, nodeRefs[neighbor])
			}
		}
	}
	return nodeRefs[node]
}
