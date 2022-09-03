package LeetCode

/**
 * 61. Rotate List
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		// 空链表或只有一个节点的链表，无论如何右移都是自身
		return head
	}
	length, p := 1, head
	for p.Next != nil {
		length++
		p = p.Next
	}
	// 此时p指向尾节点
	k = k % length
	// 找到倒数第k+1个，即正数第len-k个节点
	i, q := 1, head
	for {
		if i == length-k {
			break
		}
		i++
		q = q.Next
	}
	// 此时指向倒数第k个节点的前一个节点
	p.Next = head
	res := q.Next
	q.Next = nil
	return res
}
