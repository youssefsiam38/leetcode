/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		p1      *ListNode = list1
		p2      *ListNode = list2
		newList *ListNode
		pNew    *ListNode
	)
	if p1 != nil && (p2 == nil || p1.Val < p2.Val) {
		pNew = &ListNode{
			Val: p1.Val,
		}
		p1 = p1.Next
	} else if p2 != nil {
		pNew = &ListNode{
			Val: p2.Val,
		}
		p2 = p2.Next
	}
	newList = pNew
	for p1 != nil || p2 != nil {
		if pNew.Next != nil {
			pNew = pNew.Next
		}
		if p1 != nil && (p2 == nil || p1.Val < p2.Val) {
			pNew.Next = &ListNode{
				Val: p1.Val,
			}
			p1 = p1.Next
		} else if p2 != nil {
			pNew.Next = &ListNode{
				Val: p2.Val,
			}
			p2 = p2.Next
		}
	}
	return newList
}
