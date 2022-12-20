package two_pointers

import "sort"

func QuadrupleSumToTarget(arr []int, target int) [][]int {
	quadruples := [][]int{}
	sort.Ints(arr)
	for i := 0; i < len(arr)-4; i++ {
		for j := i + 1; j < len(arr)-2; j++ {
			left, rigth := j+1, len(arr)-1
			for left < rigth {
				if (arr[left] == arr[left+1]) && left+2 < rigth {
					left++
					continue
				}
				sum := target - arr[left] - arr[rigth]
				if sum == arr[i]+arr[j] {
					quadruples = append(quadruples, []int{
						arr[i],
						arr[j],
						arr[left],
						arr[rigth],
					})
					left++
				} else if sum > arr[i]+arr[j] {
					left++
				} else {
					rigth--
				}
			}
		}
	}
	return quadruples
}
