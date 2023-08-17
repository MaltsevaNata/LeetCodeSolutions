func closedIsland(grid [][]int) int {
    countClosed := 0
    
    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[0]); x++ {
            if grid[y][x] == 0 {
                if isClosedIsland(grid, y, x) {
                    countClosed++
                }
            } 
        }
    }
    return countClosed
}

func isClosedIsland(grid [][]int, y, x int) bool {
    queue := [][2]int{{y, x}}
    dirs := [][2]int{{1,0}, {0, 1}, {-1, 0}, {0, -1}}
    isClosed := true
    
    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        y = cell[0]
        x = cell[1]
        
        if y == 0 || y == len(grid)-1 || x== 0 || x == len(grid[0])-1 {
            isClosed = false
        }
        grid[y][x] = -1
        
        for _, dir := range dirs {
            yy := y + dir[0]
            xx := x + dir[1]
            
            if yy < 0 || yy >= len(grid) || xx < 0 || xx >= len(grid[0]) {
                continue
            }
            
            if grid[yy][xx] == 0 {
                queue = append(queue, [2]int{yy, xx})
                continue
            }
        }
    }
    return isClosed
}