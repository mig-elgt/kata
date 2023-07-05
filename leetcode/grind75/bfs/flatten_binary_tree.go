package bfs

import "fmt"

// 114. Flatten Binary Tree to Linked List
// Given the root of a binary tree, flatten the tree into a "linked list":

// The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
// The "linked list" should be in the same order as a pre-order traversal of the binary tree.

// Example 1:

// Input: root = [1,2,5,3,4,null,6]
// Output: [1,null,2,null,3,null,4,null,5,null,6]
// Example 2:

// Input: root = []
// Output: []
// Example 3:

// Input: root = [0]
// Output: [0]

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -100 <= Node.val <= 100

func flatten(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	dfs(root, &stack)
	size := len(stack)
	currNode := stack[size-1]
	currNode.Left = nil
	stack = stack[0 : size-1]
	for len(stack) > 0 {
		s := len(stack)
		node := stack[s-1]
		node.Right = currNode
		node.Left = nil
		currNode = node
		stack = stack[0 : s-1]
	}
}

func dfs(root *TreeNode, stack *[]*TreeNode) {
	if root == nil {
		return
	}
	*stack = append(*stack, root)
	dfs(root.Left, stack)
	dfs(root.Right, stack)
}

func preorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	preorder(node.Left)
	preorder(node.Right)
}
