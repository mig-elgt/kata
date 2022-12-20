package mergeintervals

import "math"

func IntervalIntersection(arr1 [][]int, arr2 [][]int) [][]int {
	result := [][]int{}
	var i, j int
	for i < len(arr1) && j < len(arr2) {
		if (arr1[i][0] >= arr2[j][0] && arr1[i][0] <= arr2[j][1]) || (arr2[j][0] >= arr1[i][0] && arr2[j][0] <= arr1[i][1]) {
			result = append(result, []int{
				int(math.Max(float64(arr1[i][0]), float64(arr2[j][0]))),
				int(math.Min(float64(arr1[i][1]), float64(arr2[j][1]))),
			})
		}
		if arr1[i][1] < arr2[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}
