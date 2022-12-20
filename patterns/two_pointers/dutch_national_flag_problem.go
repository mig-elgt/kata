package two_pointers

func DutchNationalFlagProblem(arr []int) []int {
	low, high := 0, len(arr)-1
	i := 0
	for i <= high {
		if arr[i] == 0 {
			tmp := arr[i]
			arr[i] = arr[low]
			arr[low] = tmp
			i++
			low++
		} else if arr[i] == 1 {
			i++
		} else {
			tmp := arr[i]
			arr[i] = arr[high]
			arr[high] = tmp
			high--
		}
	}
	return arr
}
