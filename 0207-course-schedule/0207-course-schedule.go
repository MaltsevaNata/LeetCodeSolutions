func canFinish(numCourses int, prerequisites [][]int) bool {
    conditions := make(map[int][]int)
    visited := make(map[int]struct{})
    for _, pr := range prerequisites {
        if _, ok := conditions[pr[0]]; ok {
            conditions[pr[0]] = append(conditions[pr[0]], pr[1])
        } else {
            conditions[pr[0]] = []int{pr[1]}
        }
    }
    stack := make(map[int]struct{})
    for i := 0; i < numCourses; i++ {
        res := dfs(i, conditions, visited, stack)
        if !res {
            return false
        }
    }
    return true
}

func dfs(node int, conditions map[int][]int, visited map[int]struct{}, stack map[int]struct{}) bool {
    if _, ok := visited[node]; ok {
        return true
    }
    conds, ok := conditions[node]
    if !ok {
        return true
    }
    if _, ok := stack[node]; ok {
        return false
    }
    
    stack[node] = struct{}{}
    for _, c := range conds {
        if _, ok := stack[c]; ok {
            return false
        }
        res := dfs(c, conditions, visited, stack)
        if !res {
            return false
        }
        delete(stack, c)
    }
    delete(stack, node)
    visited[node] = struct{}{}
    return true
}