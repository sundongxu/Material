package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	return sortedListToBSTHelper(head, 0, length-1)
}

func sortedListToBSTHelper(head *ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	left := sortedListToBSTHelper(head, start, mid-1)
	parent := &TreeNode{Val: head.Val}
	parent.Left = left
	head = head.Next
	parent.Right = sortedListToBSTHelper(head, mid+1, end)
	return parent
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head != nil && head.Next == nil {
		return &TreeNode{Val: head.Val, Left: nil, Right: nil}
	}
	middleNode, preNode := middleNodeAndPreNode(head)
	if middleNode == nil {
		return nil
	}
	if preNode != nil {
		// 将middle节点左侧链表与middle节点断开
		preNode.Next = nil
	}
	if middleNode == head {
		// 只有一个节点
		head = nil
	}
	return &TreeNode{Val: middleNode.Val, Left: sortedListToBST(head), Right: sortedListToBST(middleNode.Next)}
}

func middleNodeAndPreNode(head *ListNode) (middle *ListNode, pre *ListNode) {
	if head == nil || head.Next == nil {
		return nil, head
	}
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1, pre
}
