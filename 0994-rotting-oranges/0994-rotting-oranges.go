
func orangesRotting(grid [][]int) int { // [[2], [1]]
    var rotten [][3]int // [[0,0]]
    var fresh int // 1
    directions := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
    rows := len(grid) - 1
    cols := len(grid[0]) - 1
    
    for i, row := range grid { // 0
        for j, orange := range row { // 0
            switch orange {
            case 2: 
                var foundFresh bool
                for _, diff := range directions {
                    x := i + diff[0] // 1
                    y := j + diff[1] // 0
                    if x >= 0 && y >= 0 && x <= rows && y <= cols && grid[x][y] == 1 {
                        foundFresh = true
                        break
                    }
                }
                if foundFresh {
                    rotten = append(rotten, [3]int{i, j, 0})
                }
            case 1:
                fresh++
            }
        }
    }
    if fresh == 0 { // TODO: also check if surrounded by empty cells
        return 0
    }
    
    var freshToRotten int // [[2], [2]]
    for len(rotten) > 0 {  // [[0,0]]
        cur := rotten[0] // [0, 0, 0]
        rotten = rotten[1:] // [1,0]
        for _, diff := range directions {
            x := cur[0] + diff[0] // 1
            y := cur[1] + diff[1] // 0
            if x >= 0 && y >= 0 && x <= rows && y <= cols && grid[x][y] == 1 {
                grid[x][y] = 2
                freshToRotten++
                rotten = append(rotten, [3]int{x, y, cur[2] + 1 })
            }
        }
        if freshToRotten == fresh {
            return cur[2] + 1
        }
    }
    return -1
}