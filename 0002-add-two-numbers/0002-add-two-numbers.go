func getValOrZero(list *ListNode) (val int) {
	if list == nil {
		val = 0
	} else {
		val = list.Val
	}
	return val
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (sum *ListNode) {
	carry := 0
	sum = new(ListNode)
	cur := sum
	for l1 != nil || l2 != nil || carry > 0 {
		valSum := getValOrZero(l1) + getValOrZero(l2) + carry
		carry = valSum / 10
		cur.Next = &ListNode{Val: valSum % 10, Next: nil}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		cur = cur.Next
	}
	return sum.Next
}