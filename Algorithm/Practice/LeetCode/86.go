package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	leftDummy, rightDummy := &ListNode{}, &ListNode{}
	left, right := leftDummy, rightDummy
	p := head
	for p != nil {
		if p.Val < x {
			left.Next = &ListNode{Val: p.Val}
			left = left.Next
		} else {
			right.Next = &ListNode{Val: p.Val}
			right = right.Next
		}
		p = p.Next
	}
	left.Next = rightDummy.Next
	return leftDummy.Next
}
