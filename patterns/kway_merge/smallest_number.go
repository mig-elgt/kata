package kmerge

type ArrayItem struct {
	Value      int
	Index      int
	OwnerIndex int
}

func FindKthSmallest(lists [][]int, k int) int {
	panic("not impl")
	// minHeap := heap.NewHeap(func(a, b ArrayItem) bool { return a.Value > b.Value })
	// for i, list := range lists {
	// 	minHeap.Push(ArrayItem{list[0], 0, i})
	// }
	// var count int
	// for minHeap.Size() > 0 {
	// 	count++
	// 	item, _ := minHeap.Pop()
	// 	if count == k {
	// 		return item.Value
	// 	}
	// 	if item.Index < len(lists[item.OwnerIndex]) {
	// 		nextValue := lists[item.OwnerIndex][item.Index+1]
	// 		minHeap.Push(ArrayItem{nextValue, item.Index + 1, item.OwnerIndex})
	// 		item.Index++
	// 	}
	// }
	// return 0
}
