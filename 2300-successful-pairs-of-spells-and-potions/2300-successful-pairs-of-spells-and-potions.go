func successfulPairs(spells []int, potions []int, success int64) []int { // 7
    sort.Ints(potions) // nlogn
    
    var result = make([]int, len(spells))
    var ready = make(map[int]int)
    
    for idx, spell := range spells { // k*logn, total: klogn + nlogn = logn*(k+n)
        if val, ok := ready[spell]; ok {
            result[idx] = val
            continue
        }
        minMult := int(math.Ceil(float64(success)/float64(spell)))
        
        var left, right int
        right = len(potions) - 1 // 4
        
        if potions[right] < minMult {
            ready[spell] = 0
            continue
        }
        
        for { // logn
            mid := (left + right) / 2 // 1
            if potions[mid] >= minMult  { 
                if mid == 0 || potions[mid-1] < minMult  {
                    result[idx] = len(potions) - mid // 5 - 1 = 4
                    break
                }
                right = mid - 1 // 1
            } else {
                left = mid + 1 // 1
            }
        }
        ready[spell] = result[idx]

    }
    return result
}