func findNextGreater(indexes [][2]int) map[int]int {
    stack := []int{}
    result := make(map[int]int)
    
    for i := 0; i < len(indexes); i++ {
        result[indexes[i][0]] = -1
    }
    
    for i := 0; i < len(indexes); i++ {
        last := len(stack)-1
        for len(stack) > 0 && indexes[stack[last]][0] < indexes[i][0] {
            item := stack[last]
            stack = stack[:last]
            result[indexes[item][0]] = indexes[i][0]
            last = len(stack)-1
        }
        stack = append(stack, i)
    }
    return result
}

func oddEvenJumps(arr []int) int { // [2,3,1,1,4]
    indexes := make([][2]int, len(arr))
    for i:=0; i< len(arr); i++ {
        indexes[i] = [2]int{i, arr[i]}
    }
    sort.Slice(indexes, func(i, j int) bool {
        if indexes[i][1] == indexes[j][1] {
            return indexes[i][0] < indexes[j][0]
        }
        return indexes[i][1] < indexes[j][1] 
    })
    
    
    nextGreater := findNextGreater(indexes)

    sort.Slice(indexes, func(i, j int) bool {
        if indexes[i][1] == indexes[j][1] {
            return indexes[i][0] < indexes[j][0]
        }
        return indexes[i][1] > indexes[j][1] 
    })
    
    nextSmaller := findNextGreater(indexes)
    
    res := 1
    for i := 0; i < len(arr)-1; i++ {
        res += jump(arr, i, false, nextGreater, nextSmaller) 
    }
    return res
}


func jump(arr []int, fromIdx int, isEven bool, nextGreater, nextSmaller map[int]int) int {
    if isEven {
        next := nextSmaller[fromIdx]
        if next == -1 {
            return 0
        }
        if next == len(arr)-1 {
            return 1
        }
        return jump(arr, next, false, nextGreater, nextSmaller) 
    }

    next := nextGreater[fromIdx]
    if next == -1 {
        return 0
    }
    if next == len(arr)-1 {
        return 1
    }

    return jump(arr, next, true, nextGreater, nextSmaller)   
}