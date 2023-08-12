import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() any {
    old := *h
    item := old[len(old)-1]
    *h = old[:len(old)-1]
    return item
}

func (h MinHeap) Peek() int {
    return h[0]
}

func (h *MinHeap) Push(item any) {
    *h = append(*h, item.(int))
}


func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
    busesHeap := MinHeap{}
    for _, bus := range buses {
        busesHeap = append(busesHeap, bus)
    }
    passMap := make(map[int]struct{})
    passHeap := MinHeap{}
    for _, pass := range passengers {
        passHeap = append(passHeap, pass)
        passMap[pass] = struct{}{}
    }
    
    heap.Init(&busesHeap)
    heap.Init(&passHeap)
    
    var latestTime int // 10

    for len(busesHeap) > 0 {
        bus := heap.Pop(&busesHeap).(int) // 20
        if len(passHeap) == 0 {
            latestTime = bus
            continue
        }
        
        next := passHeap.Peek() // 18 
        num := 0
        
        for next <= bus && num < capacity && len(passHeap) > 0 {
            if _, ok := passMap[next - 1]; !ok {
                latestTime = next - 1
            }
            heap.Pop(&passHeap)
            
            num++
            if len(passHeap) > 0 {
                next = passHeap.Peek()
            }
        }
        if num < capacity && next != bus {
            if _, ok := passMap[bus]; !ok {
                latestTime = bus
            }
        }
    }
    return latestTime
}