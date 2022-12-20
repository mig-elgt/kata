package grind75

import (
	"math"
	"strconv"
)

func myAtoi(str string) int {
	sign := 1
	i := 0
	for i < len(str) {
		if str[i] == '+' || str[i] == '-' || (str[i] >= '0' && str[i] <= '9') {
			break
		}
		i++
	}
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}
	var integer string
	for i < len(str) {
		if str[i] >= '0' && str[i] <= '9' {
			if str[i] != '0' {
				integer += string(str[i])
			}
		}
		i++
	}
	number, _ := strconv.Atoi(integer)
	res := number * sign
	min := int(math.Pow(-2, 31))
	max := int(math.Pow(2, 31-1))
	if res >= min && res <= max {
		return res
	}
	return 0
}
