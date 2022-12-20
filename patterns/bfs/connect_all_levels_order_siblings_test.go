package bfs

import (
	"testing"
)

func TestConnectAllLevelOrderSiblings(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	want := "1 2 3 4 5 6 7 "
	head := ConnectAllLevelOrderSiblings(root)

	if got := getLevelString(head); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}

}
