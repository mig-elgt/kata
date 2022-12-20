package cyclicsort

func FindAllMissingNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			tmp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[tmp-1] = tmp
		} else {
			i++
		}
	}
	missingNums := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			missingNums = append(missingNums, i+1)
		}
	}
	return missingNums
}
