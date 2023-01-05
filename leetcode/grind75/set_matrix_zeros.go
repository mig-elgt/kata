package grind75

// 73. Set Matrix Zeroes
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

// You must do it in place.
// Example 1:
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]

func setZeroes(matrix [][]int) {
	zeros := [][]int{}
	rows := len(matrix)
	columns := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}
	rowZeros := map[int]int{}
	columnZeros := map[int]int{}
	for _, zero := range zeros {
		if _, ok := rowZeros[zero[0]]; !ok {
			// Convert row values to zero
			for column := 0; column < columns; column++ {
				matrix[zero[0]][column] = 0
			}
			rowZeros[zero[0]] = 1
		}
		if _, ok := columnZeros[zero[1]]; !ok {
			// Convert column values to zero
			for row := 0; row < rows; row++ {
				matrix[row][zero[1]] = 0
			}
			columnZeros[zero[1]] = 1
		}
	}
}
