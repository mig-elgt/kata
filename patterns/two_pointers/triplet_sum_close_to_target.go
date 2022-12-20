package two_pointers

import (
	"math"
	"sort"
)

func TripletSumCloseToTarget(numbers []int, target int) int {
	sort.Ints(numbers)
	distanceClose := math.MaxInt
	var sumClose int
	for i := 0; i < len(numbers)-2; i++ {
		left, rigth := i+1, len(numbers)-1
		for left < rigth {
			sum := numbers[i] + numbers[left] + numbers[rigth]
			distance := target - sum
			if distance < 0 {
				rigth--
				continue
			}
			if distanceClose > distance {
				distanceClose = distance
				sumClose = sum
			}
			if sum < target {
				left++
			} else {
				rigth--
			}
		}
	}
	return sumClose
}
