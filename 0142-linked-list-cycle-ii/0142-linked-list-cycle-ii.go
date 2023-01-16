/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})
	for head != nil {
		_, ok := seen[head]
		if !ok {
			seen[head] = struct{}{}
		} else {
			return head
		}
		head = head.Next
	}
	return nil
}