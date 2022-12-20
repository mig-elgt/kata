package bfs

import (
	"reflect"
	"testing"
)

func TestGetZigzagTraversal(t *testing.T) {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Left.Left = &TreeNode{Val: 20}
	root.Right.Left.Right = &TreeNode{Val: 17}
	root.Right.Right = &TreeNode{Val: 5}

	want := [][]int{
		{12},
		{1, 7},
		{9, 10, 5},
		{17, 20},
	}

	if got := GetZigzagTraversal(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
