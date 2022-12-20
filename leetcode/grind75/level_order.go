package grind75

// Binary Tree Level Order Traversal
// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

// Example
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	if root != nil {
		queue := list.New()
		queue.PushBack(root)
		for queue.Len() > 0 {
			size := queue.Len()
			level := []int{}
			for i := 0; i < size; i++ {
				item := queue.Front()
				queue.Remove(item)
				node := item.Value.(*TreeNode)
				level = append(level, node.Val)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
			levels = append(levels, level)
		}
	}
	return levels
}
