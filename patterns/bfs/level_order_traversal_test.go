package bfs

import (
	"reflect"
	"testing"
)

func TestGetLevelOrderTraversal(t *testing.T) {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}

	want := [][]int{
		{12},
		{7, 1},
		{9, 10, 5},
	}

	if got := GetLevelOrderTraversal(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
