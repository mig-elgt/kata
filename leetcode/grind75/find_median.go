package grind75

// Problem: Find Median from Data Stream
// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:

// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

// Example 1:

// Input
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// Output
// [null, null, null, 1.5, null, 2.0]

// Explanation
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0

import "container/heap"

type MinStreamHeap []int

func (h MinStreamHeap) Len() int           { return len(h) }
func (h MinStreamHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinStreamHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinStreamHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinStreamHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxStreamHeap []int

func (h MaxStreamHeap) Len() int           { return len(h) }
func (h MaxStreamHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxStreamHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *MaxStreamHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxStreamHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	minHeap MinStreamHeap
	maxHeap MaxStreamHeap
}

func NewMedianFinder() MedianFinder {
	minHeap := MinStreamHeap{}
	maxHeap := MaxStreamHeap{}
	heap.Init(&minHeap)
	heap.Init(&maxHeap)
	return MedianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 {
		heap.Push(&this.maxHeap, num)
		return
	}
	if num <= this.maxHeap[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}
	this.balance()
}

func (this *MedianFinder) balance() {
	if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		min := heap.Pop(&this.minHeap)
		heap.Push(&this.maxHeap, min)
	} else if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		max := heap.Pop(&this.maxHeap)
		heap.Push(&this.minHeap, max)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64(this.minHeap[0]+this.maxHeap[0]) / 2
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap[0])
	}
	return float64(this.maxHeap[0])
}
