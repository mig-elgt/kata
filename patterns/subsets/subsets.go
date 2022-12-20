package subsets

func FindSubsets(nums []int) [][]int {
	subsets := [][]int{{}}
	for _, num := range nums {
		subsetSize := len(subsets)
		for i := 0; i < subsetSize; i++ {
			subsets = append(subsets, append(subsets[i], num))
		}
	}
	return subsets
}
