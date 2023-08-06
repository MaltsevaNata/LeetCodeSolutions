type MinQueue [][]int

func (q MinQueue) Len() int {
    return len(q)
}

func (q MinQueue) Less(i, j int) bool {
    if q[i][1] == q[j][1] {
        return q[i][2] < q[j][2] // smallest index
    }
    return q[i][1] < q[j][1] // shortest processing time
}

func (q MinQueue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}

func (q *MinQueue) Push(item any) {
    *q = append(*q, item.([]int))
}

func (q *MinQueue) Pop() any {
    old := *q
    item := old[len(old)-1]
    *q = old[:len(old)-1]
    return item
}

func getOrder(tasks [][]int) []int { // [[1,2],[2,4],[3,2],[4,1]]
    result := []int{} // 0, 2, 3, 1
    
    for i, _ := range tasks { // [[1,2,0],[2,4,1],[3,2,2],[4,1,3]]
        tasks[i] = append(tasks[i], i)
    }
    
    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i][0] < tasks[j][0]
    })
    
    curTime := tasks[0][0] // 1
    
    pq := MinQueue{}
    heap.Init(&pq)
    
    idx := 0
    for idx < len(tasks) || len(pq) > 0 {
        if len(pq) == 0 && tasks[idx][0] > curTime {
            curTime = tasks[idx][0]
        }
        
        for idx < len(tasks) && tasks[idx][0] <= curTime {
            heap.Push(&pq, tasks[idx]) // [2,4,1],
            idx++ // 3
        }
        process := heap.Pop(&pq).([]int) // [4,1,3]
        result = append(result, process[2]) // index
        curTime = curTime + process[1] // 5
    }
    return result
}