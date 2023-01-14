/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) (result *ListNode) {
	result = new(ListNode)
	cur := result
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = &ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		} else if list2 == nil {
			cur.Next = &ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else if list1.Val >= list2.Val {
			cur.Next = &ListNode{Val: list2.Val, Next: nil}
			if list2.Next != nil && list2.Next.Val > list1.Val {
				cur = cur.Next
				cur.Next = &ListNode{Val: list1.Val, Next: nil}
				list1 = list1.Next
			}
			list2 = list2.Next
		} else {
			cur.Next = &ListNode{Val: list1.Val, Next: nil}
			if list1.Next != nil && list1.Next.Val > list2.Val {
				cur = cur.Next
				cur.Next = &ListNode{Val: list2.Val, Next: nil}
				list2 = list2.Next
			}
			list1 = list1.Next
		}
		cur = cur.Next
	}
	return result.Next
}