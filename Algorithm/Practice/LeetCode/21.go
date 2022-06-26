package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//[1,2,4]
//[1,3,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	p, q := list1, list2
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := head
	for p != nil && q != nil {
		node := &ListNode{
			Val:  0,
			Next: nil,
		}
		if p.Val < q.Val {
			node.Val = p.Val
			p = p.Next
		} else {
			node.Val = q.Val
			q = q.Next
		}
		tmp.Next = node
		tmp = tmp.Next
	}
	if p != nil {
		tmp.Next = p
	}
	if q != nil {
		tmp.Next = q
	}
	return head.Next
}
