package subsets

func GetPermutations(nums []int) [][]int {
	result := [][]int{}
	permutations := NewQueue([]int{})
	permutations.Push([]int{})
	for _, num := range nums {
		permutationSize := permutations.Len()
		for i := 0; i < permutationSize; i++ {
			currentPermutation := permutations.Pop()
			for j := 0; j <= len(currentPermutation); j++ {
				newPermutation := append(currentPermutation[:j], append([]int{num}, currentPermutation[j:]...)...)
				if len(newPermutation) == len(nums) {
					result = append(result, newPermutation)
				}
				permutations.Push(newPermutation)
			}
		}
	}
	return result
}
