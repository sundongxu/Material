package LeetCode

import "fmt"

// Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 先计算两个链表长度la和lb，让更长的链表先走|la-lb|步，然后两个链表一起走，如果走到相同点，则说明有交点
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	pA, pB := headA, headB

	for pA != nil {
		lenA++
		pA = pA.Next
	}

	for pB != nil {
		lenB++
		pB = pB.Next
	}

	pA, pB = headA, headB
	if lenA > lenB {
		diff := lenA - lenB
		for diff > 0 {
			pA = pA.Next
			diff--
		}
	} else {
		diff := lenB - lenA
		for diff > 0 {
			pB = pB.Next
			diff--
		}
	}

	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		} else {
			pA = pA.Next
			pB = pB.Next
		}
	}

	return nil
}

// 分别遍历两个链表，如果某个链表走到末尾了，则将另一个链表拼接到其末尾，继续遍历，如果走到相同点，说明有交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}
