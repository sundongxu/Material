package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	cur := &ListNode{Next: head}
	var lPre, l, r, rNext *ListNode
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == left {
			lPre = cur
			l = cur.Next
		} else if cur.Next.Val == right {
			r = cur.Next
			rNext = r.Next
			r.Next = nil
		}
	}
	head = reverseList(l)
	lPre.Next = head
	l.Next = rNext
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// 空链表或只有一个节点的链表，反转后就是自身
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 题目要求仅遍历一次链表
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}
	// pre已经指向第left-1个节点
	cur := pre.Next // pre之后第一个节点，将它之后的right-left个节点依次用头插法插入它之前，将原第right个节点接在pre后面
	for i := 0; i < right-left; i++ {
		tmp := cur.Next.Next
		cur.Next.Next = pre.Next
		pre.Next = cur.Next
		cur.Next = tmp
	}
	return dummy.Next
}
