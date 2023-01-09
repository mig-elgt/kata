package arrays

// 74. Search a 2D Matrix
// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		if matrix[0][0] == target {
			return true
		}
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	start, end := 0, (rows*columns)-1
	for start <= end {
		middle := start + (end-start)/2
		i := middle / columns
		j := middle % columns
		if target == matrix[i][j] {
			return true
		}
		if target < matrix[i][j] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return false
}
