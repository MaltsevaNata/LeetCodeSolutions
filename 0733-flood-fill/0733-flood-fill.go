func fillCell(image [][]int, sr int, sc int, color, origColor int) [][]int {
	if image[sr][sc] == color || image[sr][sc] != origColor {
		return image
	}
	image[sr][sc] = color
	// fill left cell
	if sr > 0 {
		image = fillCell(image, sr-1, sc, color, origColor)
	}
	// fill upper cell
	if sc > 0 {
		image = fillCell(image, sr, sc-1, color, origColor)
	}
	// fill right cell
	if sr < len(image)-1 {
		image = fillCell(image, sr+1, sc, color, origColor)
	}
	// fill down cell
	if sc < len(image[0])-1 {
		image = fillCell(image, sr, sc+1, color, origColor)
	}
	return image
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	origColor := image[sr][sc]
	return fillCell(image, sr, sc, color, origColor)
}
