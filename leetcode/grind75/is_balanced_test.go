package grind75

import "testing"

func TestIsBalancedBaseCase(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	if got := isBalanced(root); got != true {
		t.Fatalf("got %v; want true", got)
	}
}

func TestIsBalancedFailed(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}

	if got := isBalanced(root); got != false {
		t.Fatalf("got %v; want false", got)
	}
}
