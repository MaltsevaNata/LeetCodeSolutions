type Node struct {
    Start int
    End int
    Left *Node
    Right *Node
}

type MyCalendar struct {
    root *Node
}


func Constructor() MyCalendar {
    return MyCalendar{root: nil}
}

func (this *MyCalendar) insertBooking(node *Node, start, end int) bool {
    if start < node.End && end > node.Start { // overlap
        return false
    }
    if start < node.Start {
        if node.Left == nil {
            node.Left = &Node{Start: start, End: end}
            return true
        }
        return this.insertBooking(node.Left, start, end)
    }
    if node.Right == nil {
        node.Right = &Node{Start: start, End: end}
        return true
    }
    return this.insertBooking(node.Right, start, end)
}


func (this *MyCalendar) Book(start int, end int) bool {
    if this.root == nil {
        this.root = &Node{Start: start, End: end}
        return true
    }
    return this.insertBooking(this.root, start, end) // 20, 30
} // true, false, 


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */