package dfs

import (
	"reflect"
	"testing"
)

func TestGetAllPathsSum(t *testing.T) {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}

	want := [][]int{
		{12, 7, 4},
		{12, 1, 10},
	}

	if got := GetAllPathsSum(root, 23); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
