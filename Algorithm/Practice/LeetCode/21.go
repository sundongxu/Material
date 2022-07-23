package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//[1,2,4]
//[1,3,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	p, q := list1, list2
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := head
	for p != nil && q != nil {
		node := &ListNode{
			Val:  0,
			Next: nil,
		}
		if p.Val < q.Val {
			node.Val = p.Val
			p = p.Next
		} else {
			node.Val = q.Val
			q = q.Next
		}
		tmp.Next = node
		tmp = tmp.Next
	}
	if p != nil {
		tmp.Next = p
	}
	if q != nil {
		tmp.Next = q
	}
	return head.Next
}

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// dummy指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := ListNode{}
	p := &dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}
