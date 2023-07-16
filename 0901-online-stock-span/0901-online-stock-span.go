type Stack struct {
	items [][2]int
	size  int
}

func (s *Stack) Pop() [2]int {
	last := s.size - 1
	item := s.items[last]
	s.items = s.items[:last]
	s.size--
	return item
}

func (s *Stack) Push(item [2]int) {
	s.items = append(s.items, item)
	s.size++
}

type StockSpanner struct {
	stack *Stack
}

func Constructor() StockSpanner {
	return StockSpanner{stack: &Stack{}}
}

func (this *StockSpanner) Next(price int) int { // 70
	span := 1
	for this.stack.size > 0 { // {{100, 0}, {80, 1}, {70, 3}}
		item := this.stack.Pop() //
		val, days := item[0], item[1]

		if val > price {
			this.stack.Push(item)
			this.stack.Push([2]int{price, span})
			return span
		}
		span += days
	}
	this.stack.Push([2]int{price, span})
	return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */