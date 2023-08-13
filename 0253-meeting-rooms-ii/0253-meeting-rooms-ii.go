type MinHeap [][2]int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i][1] < h[j][1]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(item any) {
    *h = append(*h, item.([2]int))
}

func (h *MinHeap) Pop() any {
    old := *h
    last := len(old)-1
    item := old[last]
    *h = old[:last]
    return item
}

func (h MinHeap) Peek() any {
    return h[0]
}

func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    rooms := MinHeap{}
    rooms.Push([2]int{intervals[0][0], intervals[0][1]})
    heap.Init(&rooms)
    
    for i := 1; i < len(intervals); i++ {
        cur := intervals[i]
        if len(rooms) > 0 {
            last := rooms.Peek().([2]int)
            if cur[0] >= last[1] {
                heap.Pop(&rooms)
            } 
        }
        heap.Push(&rooms, [2]int{cur[0], cur[1]})
    }
    
    return len(rooms)
}