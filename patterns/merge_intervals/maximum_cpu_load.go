package mergeintervals

import (
	"math"
	"sort"
)

func FindMaxCPULoad(jobs [][]int) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	jobA := jobs[0]
	maxCPU := jobA[2]
	for i := 1; i < len(jobs); i++ {
		jobB := jobs[i]
		if jobB[0] <= jobA[1] {
			jobA[0] = int(math.Min(float64(jobA[0]), float64(jobB[0])))
			jobA[1] = int(math.Max(float64(jobA[1]), float64(jobB[1])))
			jobA[2] = jobA[2] + jobB[2]
			if maxCPU < jobA[2] {
				maxCPU = jobA[2]
			}
		} else if maxCPU < jobB[2] {
			maxCPU = jobB[2]
		}
	}
	return maxCPU
}
