func popFront(q *[][2]int) [2]int {
	r := (*q)[0]
	(*q)[0][0] = 0 
	(*q)[0][1] = 0
	*q = (*q)[1:len(*q)]
	return r
}

func nearestExit(maze [][]byte, entrance []int) int {
	cellsToVisit := [][2]int{{entrance[0], entrance[1]}}
	step := 0
	levels := []int{0, 0}
    
    rows := len(maze) - 1
    cols := len(maze[0]) - 1
    
	for len(cellsToVisit) > 0 {
		currentCell := popFront(&cellsToVisit)
		maze[currentCell[0]][currentCell[1]] = byte('+')

		if (currentCell[0] == 0 || currentCell[1] == 0 || currentCell[0] == rows || currentCell[1] == cols) && step > 0 {
			return step
		}

		for _, cell := range [][2]int{
			{currentCell[0] - 1, currentCell[1]}, // up
			{currentCell[0] + 1, currentCell[1]}, // down
			{currentCell[0], currentCell[1] - 1}, // left
			{currentCell[0], currentCell[1] + 1}, // right
		} {
			if cell[0] >= 0 && cell[1] >= 0 && cell[0] <= rows && cell[1] <=cols && maze[cell[0]][cell[1]] == byte('.') { // inside maze and empty
				if cell[0] == 0 || cell[1] == 0 || cell[0] == rows || cell[1] == cols {
					return step + 1
				}
				levels[step+1]++
				cellsToVisit = append(cellsToVisit, [2]int{cell[0], cell[1]})
                maze[cell[0]][cell[1]] = '+'
			}
		}
		levels[step]--
		if levels[step] <= 0 {
			levels = append(levels, 0)
			step++ 
		}
	}
	return -1
}