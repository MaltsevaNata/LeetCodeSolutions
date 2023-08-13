func countSwaps(row, other []int, target int) int {
    count := 0
    for i := 0; i < len(row); i++ {
        if row[i] != target {
            if other[i] != target {
                return -1
            }
            count++
        }
    }
    return count
}

func minDominoRotations(tops []int, bottoms []int) int {
    counts := []int{}
    
    counts = append(counts, countSwaps(tops, bottoms, tops[0]))
    counts = append(counts, countSwaps(tops, bottoms, bottoms[0]))
    counts = append(counts, countSwaps(bottoms, tops, tops[0]))
    counts = append(counts, countSwaps(bottoms, tops, bottoms[0]))

    min := 20000
    for _, val := range counts {
        if val != -1 && val < min {
            min = val
        }
    } 
    if min == 20000 {
        return -1
    }
    return min
}