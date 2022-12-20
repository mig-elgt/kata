package two_pointers

func PairWithTargetSum(numbers []int, target int) []int {
	left, rigth := 0, len(numbers)-1
	for left < rigth {
		currentSum := numbers[left] + numbers[rigth]
		if currentSum == target {
			return []int{left, rigth}
		}
		if currentSum > target {
			rigth--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}
