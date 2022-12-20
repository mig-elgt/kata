package cyclicsort

func FindCorruptPair(nums []int) []int {
	corruptPairs := []int{}
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			corruptPairs = append(corruptPairs, nums[i], i+1)
			break
		}
	}
	return corruptPairs
}
