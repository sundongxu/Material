package LeetCode

/**
 * 138. Copy List with Random Pointer
 * 描述：
 * 难度：Medium
 * 类型：Linked List
 */
// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 题目要求深拷贝
// 将原链表全部节点统统深拷贝一个，放在源节点的next
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		// 将每个节点都copy一个，并放在源节点的next位置
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil && cur.Next != nil {
		if cur.Random != nil {
			// 关键点：新节点(cur.Next)的random指向起源节点的random指向的节点(cur.Random)的下一个（也是copy出来的）
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	dummy := &Node{}
	cur = head
	newCur := dummy
	for cur != nil && cur.Next != nil {
		// 分离两个链表
		newCur.Next = cur.Next
		newCur = newCur.Next
		cur.Next = cur.Next.Next
		cur = cur.Next
	}
	return dummy.Next
}
