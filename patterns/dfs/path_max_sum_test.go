package dfs

import "testing"

func TestGetPathWithMaxSum(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Left.Left = &TreeNode{Val: 7}
	root.Right.Left.Right = &TreeNode{Val: 8}
	root.Right.Right = &TreeNode{Val: 6}
	root.Right.Right.Left = &TreeNode{Val: 9}

	want := 31
	if got := GetPathWithMaxSum(root); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}

}
