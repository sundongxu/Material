package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 能AC，但题目不允许修改节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for p != nil && p.Next != nil {
		p.Val, p.Next.Val = p.Next.Val, p.Val
		p = p.Next.Next
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next = next
		tmp := next.Next
		next.Next = cur
		cur.Next = tmp
		pre = cur
		cur = pre.Next
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for p := dummy; p != nil && p.Next != nil && p.Next.Next != nil; {
		p, p.Next, p.Next.Next, p.Next.Next.Next = p.Next, p.Next.Next, p.Next.Next.Next, p.Next
	}
	return dummy.Next
}
