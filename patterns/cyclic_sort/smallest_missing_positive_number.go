package cyclicsort

func FindSmallestMissingPositiveNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return 0
}
