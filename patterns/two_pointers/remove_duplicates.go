package two_pointers

func RemoveDuplicates(numbers []int) int {
	nextNonDuplicate := 1
	for i := 0; i < len(numbers); i++ {
		if numbers[nextNonDuplicate-1] != numbers[i] {
			numbers[nextNonDuplicate] = numbers[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}
