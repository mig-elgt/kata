package bfs

import (
	"reflect"
	"testing"
)

func TestGetRihtViewBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	want := []int{1, 3, 7}
	if got := GetRihtViewBinaryTree(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}

}
