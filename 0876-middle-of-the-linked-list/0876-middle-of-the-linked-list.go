/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var (
		p1 = head
		p2 = head
	)
	for p2 != nil {
		p2 = p2.Next
		if p2 != nil {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return p1
}
