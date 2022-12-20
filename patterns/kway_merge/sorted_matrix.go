package kmerge

type Item struct {
	rowIndex     int
	elementIndex int
}

func FindKthSmallestInSortedMatrix(matrix [][]int, k int) int {
	panic("not impl")
	// minHeap := heap.NewHeap(func(a, b Item) bool {
	// 	return matrix[a.rowIndex][a.elementIndex] > matrix[b.rowIndex][b.elementIndex]
	// })
	// for i := range matrix {
	// 	minHeap.Push(Item{i, 0})
	// }
	// var count, result int
	// for minHeap.Size() > 0 {
	// 	count++
	// 	item, _ := minHeap.Pop()
	// 	result = matrix[item.rowIndex][item.elementIndex]
	// 	if count == k {
	// 		break
	// 	}
	// 	item.elementIndex++
	// 	if item.elementIndex <= len(matrix[0]) {
	// 		minHeap.Push(item)
	// 	}
	// }
	// return result
}
