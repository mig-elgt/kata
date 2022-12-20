package twoheaps

type MedianStream struct {
	MaxHeap *Heap[int]
	MinHeap *Heap[int]
}

func NewMediaStream() MedianStream {
	return MedianStream{
		MaxHeap: NewHeap(func(a, b int) bool { return a < b }),
		MinHeap: NewHeap(func(a, b int) bool { return b < a }),
	}
}

func (ms MedianStream) InsertNum(num int) {
	n, _ := ms.MaxHeap.Peek()
	if ms.MaxHeap.Size() == 0 || n >= num {
		ms.MaxHeap.Push(num)
	} else {
		ms.MinHeap.Push(num)
	}
	if ms.MaxHeap.Size() > ms.MinHeap.Size()+1 {
		n, _ := ms.MaxHeap.Pop()
		ms.MinHeap.Push(n)
	} else if ms.MaxHeap.Size() < ms.MinHeap.Size() {
		n, _ := ms.MinHeap.Pop()
		ms.MaxHeap.Push(n)
	}
}

func (ms MedianStream) FindMedian() float64 {
	if ms.MaxHeap.Size() == ms.MinHeap.Size() {
		maxHeapNum, _ := ms.MaxHeap.Peek()
		minHeapNum, _ := ms.MinHeap.Peek()
		return float64(maxHeapNum)/2.0 + float64(minHeapNum)/2.0
	}
	n, _ := ms.MaxHeap.Peek()
	return float64(n)
}
