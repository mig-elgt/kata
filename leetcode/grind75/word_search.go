package grind75

// Word Search
// Given an m x n grid of characters board and a string word, return true if word exists in the grid.

// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

// Example 1
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true

import (
	"fmt"
)

func existWord(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, wordIndex, i, j int) bool {
	if wordIndex >= len(word) {
		return true
	}
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != word[wordIndex] {
		return false
	}
	ret := false
	board[i][j] = '#'
	rowOffsets := []int{0, 1, 0, -1}
	colOffsets := []int{1, 0, -1, 0}
	for d := 0; d < 4; d++ {
		if ret = backtrack(board, word, wordIndex+1, i+rowOffsets[d], j+colOffsets[d]); ret {
			break
		}
	}
	board[i][j] = word[wordIndex]
	return ret
}

func dfs(board [][]byte, word string, wordIndex, i, j int, exist *bool, visited map[string]int) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 {
		return
	}
	cell := fmt.Sprintf("%v%v", i, j)
	if _, isVisited := visited[cell]; isVisited {
		return
	}
	if board[i][j] != word[wordIndex] {
		return
	}
	visited[cell] = 1
	if wordIndex == len(word)-1 {
		*exist = true
		return
	}
	dfs(board, word, wordIndex+1, i, j+1, exist, visited)
	dfs(board, word, wordIndex+1, i+1, j, exist, visited)
	dfs(board, word, wordIndex+1, i, j-1, exist, visited)
	dfs(board, word, wordIndex+1, i-1, j, exist, visited)
}
