package topk

type NumberStream struct {
	k int
	// minHeap *heap.Heap[int]
}

func NewNumberStream(nums []int, k int) *NumberStream {
	panic("not impl")
	// stream := &NumberStream{
	// 	k:       k,
	// 	minHeap: heap.NewHeap(func(a, b int) bool { return a > b }),
	// }
	// for _, num := range nums {
	// 	stream.Add(num)
	// }
	// return stream
}

func (this *NumberStream) Add(num int) int {
	panic("not impl")
	// this.minHeap.Push(num)
	// if this.minHeap.Size() > this.k {
	// 	this.minHeap.Pop()
	// }
	// kth, _ := this.minHeap.Peek()
	// return kth
}
