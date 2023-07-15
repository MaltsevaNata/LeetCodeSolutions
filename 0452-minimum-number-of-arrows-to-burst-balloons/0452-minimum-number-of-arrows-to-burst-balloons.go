func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })
    
    var arrows = 1 
    start := points[0][0]
    end := points[0][1]
    
    for i := 1; i < len(points); i++ {
        if points[i][0] <= end  {
            if points[i][0] > start {
                start = points[i][0]
            }
        } else {
            arrows++
            start = points[i][0]
            end = points[i][1]
        }
    }
    return arrows
}