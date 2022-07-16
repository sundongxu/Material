package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			// 但pre不能后顺势后移至cur，因为cur是要删除的元素，cur即将指向cur.Next，而pre已经变成了cur.Next的前置节点
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
