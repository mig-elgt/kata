package grind75

// 599. Minimum Index Sum of Two Lists
/*
Given two arrays of strings list1 and list2, find the common strings with the least index sum.

A common string is a string that appeared in both list1 and list2.

A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j] then i + j should be the minimum value among all the other common strings.

Return all the common strings with the least index sum. Return the answer in any order.

Example 1:

Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
Output: ["Shogun"]
Explanation: The only common string is "Shogun".

Example 2:
Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
Output: ["Shogun"]
Explanation: The common string with the least index sum is "Shogun" with index sum = (0 + 1) = 1.

Example 3:
Input: list1 = ["happy","sad","good"], list2 = ["sad","happy","good"]
Output: ["sad","happy"]
Explanation: There are three common strings:
"happy" with index sum = (0 + 1) = 1.
"sad" with index sum = (1 + 0) = 1.
"good" with index sum = (2 + 2) = 4.
The strings with the least index sum are "sad" and "happy".

Constraints:

1 <= list1.length, list2.length <= 1000
1 <= list1[i].length, list2[i].length <= 30
list1[i] and list2[i] consist of spaces ' ' and English letters.
All the strings of list1 are unique.
All the strings of list2 are unique.
There is at least a common string between list1 and list2.

*/

import "container/heap"

type commonStr struct {
	Str string
	Sum int
}

type StrMinHeap []commonStr

func (h StrMinHeap) Len() int           { return len(h) }
func (h StrMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h StrMinHeap) Less(i, j int) bool { return h[i].Sum < h[j].Sum }
func (h StrMinHeap) Peek() commonStr {
	return h[0]
}

func (h *StrMinHeap) Push(x interface{}) {
	*h = append(*h, x.(commonStr))
}

func (h *StrMinHeap) Pop() interface{} {
	old := *h
	s := len(old)
	x := old[s-1]
	old = old[:s-1]
	*h = old
	return x
}

func findRestaurant(list1 []string, list2 []string) []string {
	commonSum := map[string]int{}
	for i, str := range list1 {
		commonSum[str] += i
	}
	minHeap := &StrMinHeap{}
	heap.Init(minHeap)
	for i, str := range list2 {
		if _, ok := commonSum[str]; ok {
			heap.Push(minHeap, commonStr{Str: str, Sum: commonSum[str] + i})
		}
	}
	var result []string
	prevSum := minHeap.Peek().Sum
	for minHeap.Len() > 0 {
		cs := heap.Pop(minHeap).(commonStr)
		if prevSum != cs.Sum {
			break
		}
		result = append(result, cs.Str)
		prevSum = cs.Sum
	}
	return result
}
