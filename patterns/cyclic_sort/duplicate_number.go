package cyclicsort

func FindDuplicateNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			tmp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[tmp-1] = tmp
		} else {
			if i > nums[i]-1 {
				return nums[i]
			} else {
				i++
			}
		}
	}
	return 0
}
