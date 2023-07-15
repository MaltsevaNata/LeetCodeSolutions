func min(i, j int) int {
    if i<j {
        return i
    }
    return j
}

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1] // sort by end time
	})
    
    end := intervals[0][1]
    var count int

    for i := 1; i < len(intervals); i++ { 
        if intervals[i][0] >= end {
            end = intervals[i][1]
        } else {
            count++
        }
    }
    return count
}