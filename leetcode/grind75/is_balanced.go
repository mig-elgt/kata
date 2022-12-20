package grind75

// Balanced Binary Tree

// Given a binary tree, determine if it is height-balanced.
// For this problem, a height-balanced binary tree is defined as:
// a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

// Example
// Input: root = [3,9,20,null,null,15,7]
// Output: true

import "math"

func isBalanced(root *TreeNode) bool {
	balanced := true
	isBalancedRecursive(root, &balanced)
	return balanced
}

func isBalancedRecursive(node *TreeNode, isBalanced *bool) int {
	if node == nil {
		return 0
	}
	leftHeight := isBalancedRecursive(node.Left, isBalanced)
	rightHeight := isBalancedRecursive(node.Right, isBalanced)
	if !*isBalanced {
		return 0
	}
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		*isBalanced = false
		return 0
	}
	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}
