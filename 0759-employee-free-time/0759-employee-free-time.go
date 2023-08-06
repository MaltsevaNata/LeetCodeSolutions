/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */
import "container/heap"
type MinHeap []*Interval

func (q MinHeap) Len() int {
    return len(q)
}

func (q MinHeap) Less(i, j int) bool {
    return q[i].Start < q[j].Start
}

func (q MinHeap) Swap (i, j int) {
    q[i], q[j] = q[j], q[i]
}

func (q *MinHeap) Push(item any) {
    *q = append(*q, item.(*Interval))
}

func (q *MinHeap) Pop() any  {
    old := *q
    item := old[len(old)-1]
    *q = old[:len(old)-1]
    return item
}

func min (i, j int) int {
    if i < j {
        return i
    }
    return j
}

func max (i, j int) int {
    if i > j {
        return i
    }
    return j
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
    h := MinHeap{}
    heap.Init(&h)
    
    merged := []*Interval{}
    result := []*Interval{}
    
    for _, intervals := range schedule {
        for _, interval := range intervals {
            heap.Push(&h, interval)
        }
    }
    merged = append(merged, heap.Pop(&h).(*Interval))
    
    for len(h) > 0 {
        interval := heap.Pop(&h).(*Interval)
        last := merged[len(merged)-1]
        
        if interval.Start < last.End && interval.End > last.Start {
            merged[len(merged)-1].Start = min(last.Start, interval.Start)
            merged[len(merged)-1].End = max(last.End, interval.End)
        } else {
            merged = append(merged, interval)
        }
    }
    
    for i := 0; i<len(merged)-1; i++ {
        if merged[i].End < merged[i+1].Start {
            result = append(result, &Interval{Start: merged[i].End, End: merged[i+1].Start})
        }
    }
    return result
}