package grind75

import "testing"

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	want := 3
	if got := maxDepth(root); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
