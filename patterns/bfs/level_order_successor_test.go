package bfs

import "testing"

func TestGetLevelOrderSuccessor(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	want := 4
	if got := GetLevelOrderSuccessor(root, 3); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}

}
