package grind75

func floodFill(image [][]int, sr, sc, color int) [][]int {
	startingPixel := image[sr][sc]
	setColorRecursive(image, sr, sc, color, startingPixel)
	return image
}

func setColorRecursive(image [][]int, sr, sc, color, startingPixel int) {
	if sr >= len(image) || sr < 0 || sc < 0 || sc >= len(image[0]) {
		return
	}
	if image[sr][sc] == startingPixel && image[sr][sc] != color {
		image[sr][sc] = color
		setColorRecursive(image, sr, sc+1, color, startingPixel)
		setColorRecursive(image, sr+1, sc, color, startingPixel)
		setColorRecursive(image, sr, sc-1, color, startingPixel)
		setColorRecursive(image, sr-1, sc, color, startingPixel)
	}
}
