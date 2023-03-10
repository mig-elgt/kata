package grind75

// 103. Binary Tree Zigzag Level Order Traversal
// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

// Example 1
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []

import "container/list"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := [][]int{}
	var currLevel int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()
		var level []int
		for i := 0; i < size; i++ {
			element := queue.Front()
			node := element.Value.(*TreeNode)
			if currLevel%2 == 0 {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			queue.Remove(element)
		}
		levels = append(levels, level)
		currLevel++
	}

	return levels
}
