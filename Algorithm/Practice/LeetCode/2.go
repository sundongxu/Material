package LeetCode

/**
 * 2. Add Two Numbers
 * 描述：
 * 难度：Medium
 * 类型：Linked List
 */
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// l1 != nil && l2 != nil
	// find the longer one list
	len1 := 0
	p1 := l1
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	len2 := 0
	p2 := l2
	for p2 != nil {
		len2++
		p2 = p2.Next
	}

	longer := l1
	shorter := l2
	if len1 < len2 {
		longer = l2
		shorter = l1
	}

	result := longer
	for longer != nil {
		shorterVal := 0
		if shorter != nil {
			shorterVal = shorter.Val
		}
		n := longer.Val + shorterVal
		if n < 10 {
			longer.Val = n
		} else {
			longer.Val = n - 10
			if longer.Next != nil {
				longer.Next.Val++
			} else {
				longer.Next = new(ListNode)
				longer.Next.Val = 1
				longer.Next.Next = nil
			}
		}
		longer = longer.Next
		if shorter != nil {
			shorter = shorter.Next
		}
	}
	return result
}
