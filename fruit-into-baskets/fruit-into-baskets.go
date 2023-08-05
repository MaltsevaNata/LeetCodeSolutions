func totalFruit(fruits []int) int { 
    if len(fruits) == 1 {
        return 1
    }
    compr := [][2]int{} 
    prev := -1 
    uq := make(map[int]struct{})
    
    for i := 0; i < len(fruits); i++ {
        if fruits[i] != prev {
            compr = append(compr, [2]int{fruits[i], 1})
            prev = fruits[i]
            uq[fruits[i]] = struct{}{}
        } else {
            compr[len(compr) - 1][1]++
        }   
    }
    if len(uq) == 2 {
        return len(fruits)
    }
    
    max := 0
    pick := make(map[int]struct{})
    left := 0
    right := 1 
    
    for right < len(compr) { 
        leftFruit := compr[left][0] 
        rightFruit := compr[right][0] 

        if _, ok := pick[leftFruit]; !ok {
            pick[leftFruit] = struct{}{}
        }

        if _, ok := pick[rightFruit]; !ok {
            if len(pick) < 2 {
                pick[rightFruit] = struct{}{}
            } else {
                count := 0
                for i := left; i < right; i++ {
                    count += compr[i][1] 
                }
                if count > max {
                    max = count
                }
                delete(pick, leftFruit)
                delete(pick, rightFruit)
                
                left++ 
                right = left+1 

                continue
            }
        }
        right++
    }
    count := 0
    for i := left; i < right; i++ {
        count += compr[i][1] 
    }
    if count > max {
        max = count
    }
    return max
}