package mergeintervals

import "sort"

func CanAttendAllAppointments(arr [][]int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	for i := 1; i < len(arr); i++ {
		if arr[i][0] < arr[i-1][1] {
			return false
		}
	}
	return true
}
