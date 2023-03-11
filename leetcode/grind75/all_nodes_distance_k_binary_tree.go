package grind75

// 863. All Nodes Distance K in Binary Tree
// Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.

// You can return the answer in any order.

// Example 1
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
// Output: [7,4,1]
// Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.

import (
	"container/list"
)

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	visited := map[int]int{}
	graph := createGraph(root)

	queue := list.New()
	queue.PushBack(target.Val)

	var distance int
	for queue.Len() > 0 {
		size := queue.Len()
		var res []int
		for i := 0; i < size; i++ {
			e := queue.Front()
			node := e.Value.(int)
			visited[node] = 1
			neighbors := graph[node]
			for j := 0; j < len(neighbors); j++ {
				n := neighbors[j]
				if _, ok := visited[n]; !ok {
					queue.PushBack(n)
				}
			}
			queue.Remove(e)
			res = append(res, node)
		}
		if distance == k {
			return res
		}
		distance++
	}
	return []int{}
}

func createGraph(node *TreeNode) map[int][]int {
	graph := map[int][]int{}

	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		e := queue.Front()
		node := e.Value.(*TreeNode)
		if node.Left != nil {
			queue.PushBack(node.Left)
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
		}
		queue.Remove(e)
	}

	return graph
}
