func firstMissingPositive(nums []int) int {
    possibleMissing := 1
    var seen = make(map[int]bool)
    
    for _, num := range nums {
        if num > 0 {
            if num == possibleMissing {
                possibleMissing ++
            }
            seen[num] = true
        }
    }
    if len(seen) == 0 {
        return possibleMissing
    }
    for {
        if _, ok := seen[possibleMissing]; ok {
            possibleMissing++
        } else {
            break
        }
    }
    return possibleMissing
}