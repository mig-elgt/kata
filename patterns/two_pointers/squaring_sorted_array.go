package two_pointers

func SquaringSortedArray(arr []int) []int {
	squares := []int{}
	left, rigth := 0, len(arr)-1
	for left <= rigth {
		leftSqr := arr[left] * arr[left]
		rigthSqr := arr[rigth] * arr[rigth]
		if leftSqr > rigthSqr {
			squares = append([]int{leftSqr}, squares...)
			left++
		} else {
			squares = append([]int{rigthSqr}, squares...)
			rigth--
		}
	}

	return squares
}
