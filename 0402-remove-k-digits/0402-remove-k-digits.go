func removeKdigits(num string, k int) string {
    if len(num) <= k {
        return "0"
    }
    
    stack := []int{}
    result := ""
    
    for _, n := range num {
        val, _ := strconv.Atoi(string(n))
        if len(stack) == 0 {
            stack = append(stack, val)
            continue
        }
        
        last := len(stack) - 1
        for k > 0 && len(stack) > 0 && val < stack[last] {
            stack = stack[:last]
            last = len(stack) -1
            k--
        }
        stack = append(stack, val)
    }
    
    if k > 0 {
        stack = stack[:len(stack)-k]
    }

    leading := true
    for _, val := range stack {
        if val == 0 && leading {
            continue
        }
        result += strconv.Itoa(val)
        leading = false
    }
    if result == "" {
        return "0"
    }
    return result
}