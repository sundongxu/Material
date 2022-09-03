package LeetCode

/**
 * 19. Remove Nth Node From End of List
 * 描述：
 * 难度：Medium
 * 类型：Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}
	if length == n {
		// delete the head
		head = nil
	}

	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := head
	prev := &ListNode{
		Val:  0,
		Next: nil,
	}
	for slow != nil {
		if fast != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next
		} else {
			// fast == nil
			// slow point to the nth node
			prev.Next = slow.Next
			// slow.Next = nil
			break
		}
	}
	return head
}
