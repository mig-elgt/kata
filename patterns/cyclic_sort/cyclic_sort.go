package cyclicsort

func Sort(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i]-1 != i {
			tmp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[tmp-1] = tmp
		} else {
			i++
		}
	}
	return nums
}
