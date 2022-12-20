package twoheaps

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap(func(a, b int) bool { return a < b })
	values := []int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}

	// Push
	for _, v := range values {
		heap.Push(v)
	}
	expected := []int{13, 12, 10, 11, 7, 4, 9, 5, 8, 6}
	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, got %v", expected, heap.data)
	}

	// Pop
	expected = []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4}
	for _, v := range expected {
		if val, ok := heap.Pop(); val != v && ok {
			t.Errorf("Expected %v, got %v", val, v)
		}
	}
}
