package grind75

// 36. Valid Sudoku
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]

// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

// Constraints:

// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	boxes := map[string]map[byte]int{}
	rows := map[int]map[byte]int{}
	columns := map[int]map[byte]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			digit := board[i][j]
			if digit != '.' {
				// Rows Validation
				if _, ok := rows[i][digit]; ok {
					return false
				} else {
					if rows[i] == nil {
						rows[i] = map[byte]int{}
					}
					rows[i][digit] = 1
				}
				// Columns Validation
				if _, ok := columns[j][digit]; ok {
					return false
				} else {
					if columns[j] == nil {
						columns[j] = map[byte]int{}
					}
					columns[j][digit] = 1
				}
				// Boxes Validation
				box := fmt.Sprintf("%v%v", i/3, j/3)
				if _, ok := boxes[box][digit]; ok {
					return false
				} else {
					if boxes[box] == nil {
						boxes[box] = map[byte]int{}
					}
					boxes[box][digit] = 1
				}
			}
		}
	}
	return true
}
