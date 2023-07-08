type SmallestInfiniteSet struct {
	removedMap map[int]bool
	curItem    int
	reAdded    MinQueue
}

type MinQueue []int

func (pq MinQueue) Len() int { return len(pq) }

func (pq MinQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq MinQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinQueue) Push(x any) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *MinQueue) Pop() any {
	old := *pq
	length := len(old)
	item := old[length-1]
	*pq = old[:length-1]
	return item
}

func (pq *MinQueue) Peek() any {
	q := *pq
	return q[len(q)-1]
}

func Constructor() SmallestInfiniteSet {
	var queue MinQueue
	heap.Init(&queue)

	removed := make(map[int]bool)
	return SmallestInfiniteSet{
		removedMap: removed,
		curItem:    1,
		reAdded:    queue,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	item := this.curItem
	var minReAdded = 1001
	if len(this.reAdded) > 0 {
		minReAdded = heap.Pop(&this.reAdded).(int)
	}
	for {
		if _, ok := this.removedMap[item]; ok {
			this.curItem++
			item = this.curItem
			continue
		}
		if minReAdded < item {
			item = minReAdded
		} else {
			this.curItem++
		}
		this.removedMap[item] = true
		return item
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.removedMap[num]; ok {
		if num < this.curItem {
			heap.Push(&this.reAdded, num)
		}
		delete(this.removedMap, num)
	}
}