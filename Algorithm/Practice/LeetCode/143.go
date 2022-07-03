package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	pre, slow, fast := &ListNode{Next: head}, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow指向中间节点，pre为中间节点前一个节点
	// 断开前后两个子链表
	pre.Next = nil
	// 反转后一半子链表
	slow = reverseList(slow)
	// 合并两个子链表
	cur := head
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = slow
		slow = slow.Next
		cur.Next.Next = tmp
		cur = tmp
	}
	cur.Next = slow
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// 空链表或只有一个节点的链表，反转后就是自身
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	cur.Next = pre
	return cur
}
