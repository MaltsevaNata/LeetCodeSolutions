func findOriginalArray(changed []int) []int {
    if len(changed) % 2 != 0 {
        return []int{}
    }
    vals := make(map[int]int)
    for _, val := range changed {
        if _, ok := vals[val]; ok {
            vals[val]++
        } else {
            vals[val] = 1
        }
    }
    result := []int{}
    
    sort.Slice(changed, func(i,j int) bool {
        return changed[i] < changed[j]
    })
    
    for _, val := range changed {
        if _, ok := vals[val]; !ok {
            continue
        }    
        if _, ok := vals[val*2]; !ok {
            return []int{}
        }   
        result = append(result, val)
        vals[val]--
        if vals[val] <= 0 {
            delete(vals, val)
        }
        vals[val*2]--
        if vals[val*2] <= 0 {
            delete(vals, val*2)
        }
    }
    
    return result
}