package LeetCode

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// 空链表或只有一个节点的链表，去重后就是自身
		return head
	}
	dummy := &ListNode{Next: head}
	pre, cur, val := dummy, head, math.MinInt
	for cur != nil && cur.Next != nil {
		val = cur.Val
		for cur.Next != nil && cur.Next.Val == val {
			cur = cur.Next
		}
		pre.Next = cur.Next
		pre = cur.Next
		cur = cur.Next.Next
	}
	return dummy.Next
}
