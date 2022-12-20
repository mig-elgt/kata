package subsets

import (
	"strconv"
	"strings"
)

func GetStringPermutationsByChangingCase(str string) []string {
	results := []string{}
	permutations := NewQueue("foobar")
	permutations.Push("")
	for i := 0; i < len(str); i++ {
		permutationsSize := permutations.Len()
		for j := 0; j < permutationsSize; j++ {
			currentPermutation := permutations.Pop()
			newPermutationLowerCase := currentPermutation + string(str[i])
			permutations.Push(newPermutationLowerCase)
			if len(newPermutationLowerCase) == len(str) {
				results = append(results, newPermutationLowerCase)
			}
			if _, err := strconv.Atoi(string(str[i])); err != nil {
				newPermutationUpperCase := currentPermutation + strings.ToUpper(string(str[i]))
				permutations.Push(newPermutationUpperCase)
				if len(newPermutationUpperCase) == len(str) {
					results = append(results, newPermutationUpperCase)
				}
			}
		}
	}
	return results
}
