package grind75

//Invert Binary Tree
// Given the root of a binary tree, invert the tree, and return its root.
//
// Example
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invertTree(root.Left)
		invertTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
