package two_pointers

func SubArraysWithProductLessThanATarget(arr []int, target int) [][]int {
	subs := [][]int{}
	for i := 0; i < len(arr); i++ {
		next := i
		product := 1
		for next < len(arr) {
			product *= arr[next]
			if product < target {
				subs = append(subs, arr[i:next+1])
			} else {
				break
			}
			next++
		}
	}
	return subs
}
