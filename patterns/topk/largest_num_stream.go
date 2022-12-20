package topk

import (
	"gitlab.com/migel/kata/heap"
)

type NumberStream struct {
	k       int
	minHeap *heap.Heap[int]
}

func NewNumberStream(nums []int, k int) *NumberStream {
	stream := &NumberStream{
		k:       k,
		minHeap: heap.NewHeap(func(a, b int) bool { return a > b }),
	}
	for _, num := range nums {
		stream.Add(num)
	}
	return stream
}

func (this *NumberStream) Add(num int) int {
	this.minHeap.Push(num)
	if this.minHeap.Size() > this.k {
		this.minHeap.Pop()
	}
	kth, _ := this.minHeap.Peek()
	return kth
}
