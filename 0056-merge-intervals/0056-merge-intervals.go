func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}


func merge(intervals [][]int) [][]int {
    if len(intervals) == 1 {
        return intervals
    }
    
    sort.Slice(intervals, func (i, j int) bool  {
        return intervals[i][0] < intervals[j][0] // by start time
    })
    
    result := [][]int{intervals[0]}

    
    for i := 1; i < len(intervals); i++ {
        last := len(result)-1

        if intervals[i][0] <= result[last][1] && intervals[i][1] >= result[last][0] { // overlap
            start := min(intervals[i][0], result[last][0])
            end := max(intervals[i][1], result[last][1])
            result[last][0] = start
            result[last][1] = end
            continue
        }
        result = append(result, intervals[i])
    }
    return result
}