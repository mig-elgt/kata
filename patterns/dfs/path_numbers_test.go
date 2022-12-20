package dfs

import "testing"

func TestGetSumPathNumbers(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 9}

	want := 408
	if got := GetSumPathNumbers(root); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
