package grind75

// Problem: Top K Frequent Words
// Given an array of strings words and an integer k, return the k most frequent strings.
// Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.

// Example 1:
// Input: words = ["i","love","leetcode","i","love","coding"], k = 2
// Output: ["i","love"]
// Explanation: "i" and "love" are the two most frequent words.
// Note that "i" comes before "love" due to a lower alphabetical order.

import (
	"container/heap"
)

type WordFrequency struct {
	word      string
	frequency int
}

type MaxHeapFrequent []WordFrequency

func (h MaxHeapFrequent) Len() int { return len(h) }
func (h MaxHeapFrequent) Less(i, j int) bool {
	if h[i].frequency == h[j].frequency {
		return h[i].word < h[j].word
	}
	return h[i].frequency > h[j].frequency
}
func (h MaxHeapFrequent) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapFrequent) Push(x interface{}) {
	*h = append(*h, x.(WordFrequency))
}

func (h *MaxHeapFrequent) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordsFrequency := map[string]int{}
	for _, w := range words {
		wordsFrequency[w]++
	}
	maxHeap := &MaxHeapFrequent{}
	heap.Init(maxHeap)
	for word, freq := range wordsFrequency {
		heap.Push(maxHeap, WordFrequency{word, freq})
	}
	result := []string{}
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(maxHeap).(WordFrequency).word)
	}
	return result
}
