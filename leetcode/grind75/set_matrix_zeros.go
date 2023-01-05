package grind75

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
	for _, zero := range zeros {
		// Convert row values to zero
		for column := 0; column < columns; column++ {
			matrix[column][zero[0]] = 0
		}
		// Convert column values to zero
		for row := 0; row < rows; row++ {
			matrix[zero[1]][row] = 0
		}
	}
}
