package LeetCode

// 25. Reverse Nodes in k-Group
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// k个节点为一组，逆转每组节点，多出来<k个节点保持不变
// 题目要求空间复杂度O(1)，即常数空间
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	// 1 -> 2 -> 3，k=2，故head指向1，node指向3
	// node指向3 => node实际已经是下一组的头节点了
	newHead := reverse(head, node) // => 2 —> 1 -> 3
	// head永远指向1 => head已经指向本组节点的末尾节点了
	head.Next = reverseKGroup(node, k) // 将本组和下一组相连
	return newHead                     // 返回本组头指针
}

// last保持不变，first在方法结束时会=last
func reverse(first *ListNode, last *ListNode) *ListNode {
	p := last
	for first != last {
		tmp := first.Next
		first.Next = p
		p = first
		first = tmp
	}
	return p
}
