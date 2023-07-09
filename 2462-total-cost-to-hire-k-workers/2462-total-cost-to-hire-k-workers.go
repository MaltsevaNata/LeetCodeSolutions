type Candidate struct {
	Cost int
	Idx  int
}

type MinHeap []*Candidate

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
    if h[i].Cost == h[j].Cost {
        return h[i].Idx < h[j].Idx
    }
	return h[i].Cost < h[j].Cost
}

func (h *MinHeap) Pop() any {
	old := *h
	length := len(old) - 1
	item := old[length]
	old[length] = nil
	*h = old[:length]
	return item
}

func (h MinHeap) Peek() any {
	return h[0]
}

func (h *MinHeap) Push(item any) {
	*h = append(*h, item.(*Candidate))
}

func totalCost(costs []int, k int, candidates int) int64 { // n -costs, k, c
    if len(costs) == 1 {
		return int64(costs[0])
	}
	var h MinHeap
	var totalCost int64 // O(n + n + k * n) = O(n * k)
	var costsQueue []int
    
    head := candidates // 4
    tail := len(costs) - candidates - 1 //  4
    
    if tail < head {
        tail = head - 1
    }

	for i := head; i <= tail; i++ { // O(n)
        costsQueue = append(costsQueue,  costs[i]) // [7]
	}


	for idx := 0; idx < head; idx++ { // O(c)
		h = append(h, &Candidate{Cost: costs[idx], Idx: idx}) // [17, 12, 10, 7, 11, 20, 8]
	}
    for idx := tail + 1; idx < len(costs); idx++ { // O(n-c)
        h = append(h, &Candidate{Cost: costs[idx], Idx: idx})
    }
    
	heap.Init(&h) // O(n)

	for i := 0; i < k; i++ { // O(k * (logn + n)
        worker := heap.Pop(&h).(*Candidate) // {7, 4}
        if len(costsQueue) > 0 {
            var next int
            if worker.Idx < head { // 3 < 4
                next = costsQueue[0] // 4
                costsQueue = costsQueue[1:] // []
                heap.Push(&h, &Candidate{Cost: next, Idx: head}) // {7, 4}
                head ++ // 5
            } else {
                length := len(costsQueue)-1
                next = costsQueue[length]
                costsQueue = costsQueue[:length]
                heap.Push(&h, &Candidate{Cost: next, Idx: tail})
                tail--
            }
        }
        
		totalCost += int64(worker.Cost) // 11
	}
	return totalCost
}