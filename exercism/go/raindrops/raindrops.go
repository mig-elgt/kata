package raindrops

import "fmt"

func Convert(number int) string {
	var factors string
	if number%3 == 0 {
		factors += "Pling"
	}
	if number%5 == 0 {
		factors += "Plang"
	}
	if number%7 == 0 {
		factors += "Plong"
	}
	if factors == "" {
		factors = fmt.Sprintf("%d", number)
	}
	return factors
}
