package bfs

import (
	"reflect"
	"testing"
)

func TestCalculateLevelAverage(t *testing.T) {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 9}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}

	want := []float64{12.0, 4.0, 6.5}

	if got := CalculateLevelAverage(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}

}
