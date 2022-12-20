package mergeintervals

import (
	"math"
	"sort"
)

func FindEmployeeFreeTime(schedule [][]int) [][]int {
	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i][0] < schedule[j][0]
	})
	freeTimes := [][]int{}
	a := schedule[0]
	for i := 1; i < len(schedule); i++ {
		if schedule[i][0] <= a[1] {
			a[0] = int(math.Min(float64(a[0]), float64(schedule[i][0])))
			a[1] = int(math.Max(float64(a[1]), float64(schedule[i][1])))
		} else {
			freeTimes = append(freeTimes, []int{
				a[1],
				schedule[i][0],
			})
			a = schedule[i]
		}
	}
	return freeTimes
}
