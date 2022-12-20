package grind75

import "fmt"

// Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).

// A valid BST is defined as follows:

// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	aux := t
	return t.stringRec(aux)
}

func (t TreeNode) stringRec(node *TreeNode) string {
	if node == nil {
		return ""
	}
	return t.stringRec(node.Left) + fmt.Sprintf("%v", node.Val) + t.stringRec(node.Right)
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(root, low, high *TreeNode) bool {
	if root == nil {
		return true
	}
	if (low != nil && root.Val <= low.Val) || (high != nil && root.Val >= high.Val) {
		return false
	}
	return validate(root.Right, root, high) && validate(root.Left, low, root)
}
