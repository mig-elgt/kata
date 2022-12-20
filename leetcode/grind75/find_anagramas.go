package grind75

import "reflect"

func findAnagrams(s string, p string) []int {
	results := []int{}
	pCount := map[rune]int{}
	for _, chr := range p {
		pCount[chr]++
	}
	sCount := map[rune]int{}
	left, rigth := 0, 0
	for rigth < len(s) {
		chr := rune(s[rigth])
		sCount[chr]++
		if rigth-left == len(p)-1 {
			if reflect.DeepEqual(pCount, sCount) {
				results = append(results, left)
			}
			leftChr := rune(s[left])
			sCount[leftChr]--
			if sCount[leftChr] == 0 {
				delete(sCount, leftChr)
			}
			left++
		}
		rigth++
	}
	return results
}
