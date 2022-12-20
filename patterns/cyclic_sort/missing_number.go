package cyclicsort

func FindMissingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] < len(nums) && nums[i] != nums[nums[i]] {
			tmp := nums[i]
			nums[i] = nums[nums[i]]
			nums[tmp] = tmp
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}
