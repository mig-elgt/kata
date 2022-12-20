package grind75

// Binary Tree Right Side View

// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

// Example 1
// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]

import (
	"container/list"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodes := []int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		level := []int{}
		for i := 0; i < size; i++ {
			element := queue.Front()
			node := queue.Remove(element).(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		nodes = append(nodes, level[len(level)-1])
	}
	return nodes
}
