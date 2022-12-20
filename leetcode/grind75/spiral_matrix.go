package grind75

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	rows := len(matrix)
	columns := len(matrix[0])
	up, left, right, down := 0, 0, columns-1, rows-1
	for len(result) < rows*columns {
		for col := left; col <= right; col++ {
			result = append(result, matrix[up][col])
		}
		for row := up + 1; row <= down; row++ {
			result = append(result, matrix[row][right])
		}
		if up != down {
			for col := right - 1; col >= left; col-- {
				result = append(result, matrix[down][col])
			}
		}
		if left != right {
			for row := down - 1; row > up; row-- {
				result = append(result, matrix[row][left])
			}
		}
		left++
		right--
		up++
		down--
	}
	return result
}
