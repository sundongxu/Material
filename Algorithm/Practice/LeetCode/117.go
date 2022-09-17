package LeetCode

/**
 * 117. Populating Next Right Pointers in Each Node II
 * 描述：
 * 难度：Medium
 * 类型：Tree & Recursion
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// 普通二叉树，要求常数空间
// 法一：递归
func connect(root *Node) *Node {
	connectHelper(root)
	return root
}

func connectHelper(root *Node) {
	if root == nil {
		return
	}
	dummy := &Node{} // 实际上永远指向root最左孩子的左边位置，dummy.Next就是root的最左孩子
	cur, prev := root, dummy
	for cur != nil {
		if cur.Left != nil {
			prev.Next = cur.Left
			prev = prev.Next
		}
		if cur.Right != nil {
			prev.Next = cur.Right
			prev = prev.Next
		}
		cur = cur.Next
	}
	// root、root的左右孩子都处理完了，处理下一层，即root的最左孩子，即dummy.Next
	connectHelper(dummy.Next)
}

// 法二：迭代
func connect(root *Node) *Node {
	cur := root
	for cur != nil {
		prev, next := &Node{}, &Node{}
		next = nil // golang不支持直接将对象声明为具体类型的nil
		for cur != nil {
			if next == nil {
				if cur.Left != nil {
					next = cur.Left
				} else {
					next = cur.Right
				}
			}
			if cur.Left != nil {
				prev.Next = cur.Left
				prev = prev.Next
			}
			if cur.Right != nil {
				prev.Next = cur.Right
				prev = prev.Next
			}
			cur = cur.Next
		}
		cur = next
	}
	return root
}
