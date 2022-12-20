package twoheaps

func FindSlidingWindowMedian(nums []int, k int) []float64 {
	medianStream := NewMediaStream()
	medians := []float64{}
	next := 0
	for i := 0; i < len(nums); i++ {
		medianStream.InsertNum(nums[i])
		if i-next >= k-1 {
			medians = append(medians, medianStream.FindMedian())
			medianStream.MaxHeap.Pop()
			next++
		}
	}
	return medians
}
