package dfs

import "testing"

func TestCountPathsForSum(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}

	want := 3
	target := 12

	if got := CountPathsForSum(root, target); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
