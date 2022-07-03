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
// 每次找链表的中间节点作为根，其左右部分链表作为根的左右子树
// 法一：用快慢指针法来找链表的中间节点
//func sortedListToBST(head *ListNode) *TreeNode {
//	if head == nil {
//		return nil
//	}
//	return sortedListToBSTHelper(head, nil)
//}
//
//func sortedListToBSTHelper(head, tail *ListNode) *TreeNode {
//	if head == tail {
//		return nil
//	}
//	// 快慢指针，慢指针走1步，快指针走2步，快指针走到最后，慢指针指向中间节点
//	slow, fast := head, head
//	for fast != tail && fast.Next != tail {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	// 以链表中间节点为根，以左右链表部分递归构造其左右子树
//	root := &TreeNode{Val: slow.Val}
//	root.Left = sortedListToBSTHelper(head, slow)
//	root.Right = sortedListToBSTHelper(slow.Next, tail)
//	return root
//}

// 法二：通过二分来找链表的中间节点
func sortedListToBST(head *ListNode) *TreeNode {
	cur, length := head, 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	return sortedListToBSTHelper(head, length)
}

func sortedListToBSTHelper(head *ListNode, length int) *TreeNode {
	// 以下收敛条件必须用length而不能用head==nil做判断，是因为head永远不会是nil，head.Next也不是nil，因为链表没有被断开
	// length才是左右子树真正的边界
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{Val: head.Val}
	}
	middle := findNthNode(head, length/2+1)
	root := &TreeNode{Val: middle.Val}
	root.Left = sortedListToBSTHelper(head, length/2)
	root.Right = sortedListToBSTHelper(middle.Next, (length-1)/2)
	return root
}

func findNthNode(head *ListNode, n int) *ListNode {
	cur := head
	for {
		n--
		if n <= 0 {
			break
		}
		cur = cur.Next
	}
	return cur
}
