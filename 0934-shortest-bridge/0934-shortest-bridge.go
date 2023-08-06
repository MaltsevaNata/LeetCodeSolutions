func shortestBridge(grid [][]int) int {
    visited := make([][]bool, len(grid))
    for row, cols := range grid {
        visited[row] = make([]bool, len(cols))
    }
    
    island := [][3]int{}
    
    for row, cols := range grid {
        for col, val := range cols {
            if !visited[row][col] {
                if val == 1 {
                    island = findIsland(row, col, visited, grid)
                    break
                }
                visited[row][col] = true
            }
        }
        if len(island) != 0 {
            break
        }
    }
    
    bridge := findBridge(island, grid)
    return bridge
}

func findBridge(queue [][3]int, grid [][]int) int {
    min := 100
    visited := make(map[[2]int]struct{})
    
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        
        y := item[0]
        x := item[1]

        if grid[y][x] == 1 {
            if item[2] < min {
                min = item[2]
            }
            continue
        }
        coord := [2]int{y, x}
        visited[coord] = struct{}{}
        if grid[y][x] == 0 {
            item[2]++
        }
        
        dirs := [][2]int{{0,1}, {0,-1}, {1, 0}, {-1,0}}
        
        for _, dir := range dirs {
            yy := y + dir[0]
            xx := x + dir[1]
            if yy >= 0 && xx >= 0 && yy < len(grid) && xx < len(grid[0]) {
                if grid[yy][xx] == 1 {
                    return item[2]
                }
                coord := [2]int{yy, xx}
                if _, ok := visited[coord]; !ok {
                    queue = append(queue, [3]int{yy, xx, item[2]})
                    visited[coord] = struct{}{}
                }
            }
        }
    }
    return min
}

func findIsland(row, col int, visited [][]bool, grid [][]int) [][3]int {
    result := [][3]int{}
    
    coord := [2]int{row, col}
    stack := [][2]int{coord}
    
    for len(stack) > 0 {
        last := len(stack) - 1
        item := stack[last]
        stack = stack[:last]
        
        y := item[0]
        x := item[1]
        if !visited[y][x] && grid[y][x] == 1 {
            result = append(result, [3]int{y, x, 0})
            // right
            if x < len(grid[0])-1 {
                stack = append(stack, [2]int{y, x+1})
            }
            // left
            if x > 0 {
                stack = append(stack, [2]int{y, x-1})
            }
            // up
            if y > 0 {
                stack = append(stack, [2]int{y-1, x})
            }
            // down
            if y < len(grid) - 1 {
                stack = append(stack, [2]int{y+1, x})
            }
            grid[y][x] = 2
        }
        visited[y][x] = true
    }
    return result
}