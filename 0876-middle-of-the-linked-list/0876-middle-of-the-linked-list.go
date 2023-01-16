/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var values []*ListNode
	for head != nil {
		values = append(values, head)
		head = head.Next
	}
	return values[len(values)/2]
}