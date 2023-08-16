func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
    if jug1Capacity+jug2Capacity < targetCapacity {
        return false
    }
    
    seen := make(map[int]struct{})
    seen[0] = struct{}{}
    ops := []int{jug1Capacity, jug2Capacity, -jug1Capacity, -jug2Capacity}
    queue := []int{0}
    
    for len(queue) > 0 {
        total := queue[0]
        queue = queue[1:]
        
        for _, op := range ops {
            val := total + op
            if val == targetCapacity {
                return true
            }
            if val < 0 || val > jug1Capacity+jug2Capacity {
                continue
            }
            if _, ok := seen[val]; ok {
                continue
            }
            seen[val] = struct{}{}
            queue = append(queue, val)
        }
    }
    return false
}