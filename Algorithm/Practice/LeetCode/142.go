package LeetCode

/**
 * 142. Linked List Cycle II
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
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow2 := head
			for slow != slow2 {
				slow = slow.Next
				slow2 = slow2.Next
			}
			return slow
		}
	}
	return nil
}
