/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) (result *ListNode) {
	var prev *ListNode = nil
	for head != nil {
		result = &ListNode{Val: head.Val, Next: prev}
		prev = result
		head = head.Next
	}
	return result
}