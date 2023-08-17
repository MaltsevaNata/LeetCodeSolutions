func updateMatrix(mat [][]int) [][]int {
    result := make([][]int, len(mat))
    for i := 0; i < len(mat); i++ {
        result[i] = make([]int, len(mat[0]))
        for j := 0; j < len(mat[0]); j++{
            if mat[i][j] == 0 {
                result[i][j] = 0
            } else {
                result[i][j] = -1
            }
        }
    }
    
    dirsDown := [][2]int{{1, 0}, {0, 1}}
    dirsUp := [][2]int{{-1, 0}, {0, -1}}
    
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j ++ {
            fillCell(i, j, result, dirsUp)
        }
    }
    for i := len(mat)-1; i >= 0; i-- {
        for j := len(mat[0])-1; j >= 0; j -- {
            fillCell(i, j, result, dirsDown)
        }
    }
    return result
}

func fillCell(y, x int, result [][]int, dirs [][2]int) {
    if result[y][x] == 0 {
        return
    }
    min := 10000
    if result[y][x] != -1 {
        min = result[y][x]
    }
    for _, dir := range dirs {
        yy := y + dir[0]
        xx := x + dir[1]

        if yy < 0 || yy >= len(result) || xx < 0 || xx >= len(result[0]) {
            continue
        }
        if result[yy][xx] != -1 {
            if result[yy][xx] + 1 < min {
                min = result[yy][xx] + 1
            }
            continue
        }
    }
    result[y][x] = min
}
