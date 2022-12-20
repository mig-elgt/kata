package two_pointers

import (
	"fmt"
	"sort"
)

func TripletSumToZero(numbers []int) [][]int {
	sort.Ints(numbers)
	sumFrecuency := map[string]int{}
	tripletsSum := [][]int{}
	for i := 0; numbers[i] < 0; i++ {
		target := numbers[i]
		left, rigth := i+1, len(numbers)-1
		for left < rigth {
			sum := target + numbers[left] + numbers[rigth]
			if sum == 0 {
				triplet := []int{numbers[i], numbers[left], numbers[rigth]}
				tripletKey := fmt.Sprintf("%v", triplet)
				if _, ok := sumFrecuency[tripletKey]; !ok {
					tripletsSum = append(tripletsSum, triplet)
					sumFrecuency[tripletKey] = 1
				}
				rigth--
			} else if sum > 0 {
				rigth--
			} else {
				left++
			}
		}
	}
	return tripletsSum
}
