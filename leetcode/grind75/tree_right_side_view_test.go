package grind75

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}

	want := []int{1, 3, 4}
	if got := rightSideView(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestRightSideViewNotRightNodes(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	want := []int{1, 2}
	if got := rightSideView(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
