package LeetCode

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
	pre, cur := dummy, head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next // cur循环结束时指向重复节点的最后一个节点
		}
		pre.Next = cur
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}
