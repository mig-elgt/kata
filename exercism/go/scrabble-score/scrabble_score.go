package scrabble

import "strings"

var scores = map[int][]rune{
	1:  {'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't'},
	2:  {'d', 'g'},
	3:  {'b', 'c', 'm', 'p'},
	4:  {'f', 'h', 'v', 'w', 'y'},
	5:  {'k'},
	8:  {'j', 'x'},
	10: {'q', 'z'},
}

func Score(word string) int {
	if word == "" {
		return 0
	}
	var score int
	chars := strings.ToLower(word)
	for _, c := range chars {
		for k, s := range scores {
			for _, letter := range s {
				if letter == c {
					score += k
				}
			}
		}
	}
	return score
}
