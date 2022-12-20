package subsets

import "strings"

func GenerateValidParentheses(num int) []string {
	results := []string{}
	permutations := NewQueue("")
	permutations.Push("")
	for permutations.Len() > 0 {
		currentPermutation := permutations.Pop()
		if currentPermutation == "" {
			permutations.Push("(")
			continue
		}
		openCount := strings.Count(currentPermutation, "(")
		closeCount := strings.Count(currentPermutation, ")")
		if openCount < num {
			newPermutation := currentPermutation + "("
			permutations.Push(newPermutation)
			if len(newPermutation) == num*2 {
				results = append(results, newPermutation)
			}
		}
		if closeCount < openCount {
			newPermutation := currentPermutation + ")"
			permutations.Push(newPermutation)
			if len(newPermutation) == num*2 {
				results = append(results, newPermutation)
			}
		}
	}
	return results
}
