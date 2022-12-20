package islands

type pixel struct {
	isVisited   bool
	isRemoved   bool
	value, i, j int
}

func removeIsland(matrix [][]int) [][]int {
	islandMatrix := createIslandMatrix(matrix)
	borderPixels := findBlackBorderPixels(matrix)
	findPixelsToRemove(islandMatrix, borderPixels)
	removeBlackPixels(matrix, islandMatrix)
	return matrix
}

func createIslandMatrix(matrix [][]int) [][]*pixel {
	islandMatrix := [][]*pixel{}
	for i := 0; i < len(matrix); i++ {
		row := make([]*pixel, len(matrix))
		islandMatrix = append(islandMatrix, row)
		for j := 0; j < len(matrix); j++ {
			if i == 0 || i == len(matrix)-1 || j == 0 || j == len(matrix)-1 {
				islandMatrix[i][j] = &pixel{
					isVisited: false,
					isRemoved: true,
					i:         i,
					j:         j,
					value:     0,
				}
			} else {
				islandMatrix[i][j] = &pixel{
					isVisited: false,
					isRemoved: true,
					i:         i,
					j:         j,
					value:     matrix[i][j],
				}

			}
		}
	}
	return islandMatrix
}

func findBlackBorderPixels(matrix [][]int) []*pixel {
	blackBorderPixels := []*pixel{}
	for i := 0; i < len(matrix); i++ {
		if i != 0 && i != len(matrix)-1 {
			if matrix[0][i] == 1 {
				blackBorderPixels = append(blackBorderPixels, &pixel{i: 0, j: i})
			}
			bb := len(matrix) - 1
			if matrix[bb][len(matrix)-1-i] == 1 {
				blackBorderPixels = append(blackBorderPixels, &pixel{i: bb, j: len(matrix) - 1 - i})
			}
			if matrix[bb-i][0] == 1 {
				blackBorderPixels = append(blackBorderPixels, &pixel{i: bb - i, j: 0})
			}
			if matrix[bb-i][bb] == 1 {
				blackBorderPixels = append(blackBorderPixels, &pixel{i: bb - i, j: bb})
			}
		}
	}
	return blackBorderPixels
}

func findPixelsToRemove(matrix [][]*pixel, borderPixels []*pixel) {
	for _, p := range borderPixels {
		matrix[p.i][p.j].isVisited = true
		if p.j == len(matrix)-1 {
			findAndSetPath(matrix[p.i][p.j-1], matrix)
		} else if p.j == 0 {
			findAndSetPath(matrix[p.i][p.j+1], matrix)
		} else if p.i == 0 {
			findAndSetPath(matrix[p.i+1][p.j], matrix)
		} else if p.i == len(matrix)-1 {
			findAndSetPath(matrix[p.i-1][p.j], matrix)
		}
	}
}

func findAndSetPath(p *pixel, islandMatrix [][]*pixel) {
	if !p.isVisited && p.value == 1 {
		p.isVisited = true
		p.isRemoved = false
		findAndSetPath(islandMatrix[p.i][p.j-1], islandMatrix)
		findAndSetPath(islandMatrix[p.i-1][p.j], islandMatrix)
		findAndSetPath(islandMatrix[p.i][p.j+1], islandMatrix)
		findAndSetPath(islandMatrix[p.i+1][p.j], islandMatrix)
	}
}

func removeBlackPixels(matrix [][]int, islandMatrix [][]*pixel) {
	for i := 1; i < len(islandMatrix)-1; i++ {
		for j := 1; j < len(islandMatrix)-1; j++ {
			if islandMatrix[i][j].isRemoved {
				matrix[i][j] = 0
			}
		}
	}
}
