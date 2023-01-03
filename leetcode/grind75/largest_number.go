package grind75

// 179. Largest Number
// Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

// Since the result may be very large, so you need to return a string instead of an integer.

// Example 1:

// Input: nums = [10,2]
// Output: "210"
// Example 2:

// Input: nums = [3,30,34,5,9]
// Output: "9534330"

import (
	"fmt"
	"math"
	"sort"
)

type numFactor struct {
	Num    int
	Factor int
}

func largestNumber(nums []int) string {
	sort.Ints(nums)
	max := nums[len(nums)-1]
	maxTotalDigits := countDigits(max)
	nf := []*numFactor{}
	for _, num := range nums {
		nf = append(nf, &numFactor{
			Num:    num,
			Factor: maxTotalDigits - countDigits(num),
		})
	}
	for _, item := range nf {
		newFact := math.Pow(10, float64(item.Factor))
		item.Num = item.Num * int(newFact)
	}
	sort.Slice(nf, func(i, j int) bool {
		return nf[i].Num > nf[j].Num
	})
	var result string
	for _, item := range nf {
		if item.Factor > 0 {
			result += fmt.Sprintf("%v", int(float64(item.Num)/float64(math.Pow(10, float64(item.Factor)))))
		} else {
			result += fmt.Sprintf("%v", float64(item.Num))
		}
	}
	return result
}

func countDigits(num int) int {
	var count int
	n := float64(num)
	for n >= 10 {
		count++
		n /= 10.0
	}
	return count + 1
}
