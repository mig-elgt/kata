package grind75

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	if got := cloneGraph(node1); !reflect.DeepEqual(got, node1) {
		t.Fatalf("got %v; want %v", got, node1)
	}

}

func TestCloneGraphNotNeighbors(t *testing.T) {
	node := &Node{Val: 1}
	if got := cloneGraph(node); !reflect.DeepEqual(node, got) {
		t.Fatalf("got %v; want %v", got, node)
	}
}

func TestCloneGraphEmpty(t *testing.T) {
	var node *Node
	if got := cloneGraph(node); !reflect.DeepEqual(node, got) {
		t.Fatalf("got %v; want %v", got, nil)
	}
}
