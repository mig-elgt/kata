package bfs

import (
	"fmt"
	"testing"
)

func TestConnectLevelOrderSiblings(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	want := "4 5 6 7 "
	head := ConnectLevelOrderSiblings(root)

	if got := getLevelString(head.Left.Left); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func getLevelString(node *TreeNode) string {
	var result string
	current := node
	for current != nil {
		result += fmt.Sprintf("%v ", current.Val)
		current = current.Next
	}
	return result
}
