package cyclicsort

func FindFirstKMissingPositiveNumbers(nums []int, k int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	bigger := nums[len(nums)-1]
	missingNums := []int{}
	for i := 0; i < len(nums); i++ {
		if bigger < nums[i] {
			bigger = nums[i]
		}
		if len(missingNums) == k {
			return missingNums
		}
		if nums[i] != i+1 {
			missingNums = append(missingNums, i+1)
		}
	}
	total := bigger + k - len(missingNums)
	for i := bigger + 1; i <= total; i++ {
		missingNums = append(missingNums, i)
	}
	return missingNums
}
