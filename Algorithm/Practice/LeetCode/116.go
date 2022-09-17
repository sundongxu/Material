package LeetCode

/**
 * 116. Populating Next Right Pointers in Each Node
 * 描述：
 * 难度：Medium
 * 类型：Tree & Recursion
 */
// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 完美二叉树
// 递归1
func connect(root *Node) *Node {
	return connectHelper(root, nil)
}

func connectHelper(root *Node, sibling *Node) *Node {
	if root == nil {
		return nil
	} else {
		root.Next = sibling
	}
	connectHelper(root.Left, root.Right)
	if sibling != nil {
		connectHelper(root.Right, sibling.Left)
	} else {
		connectHelper(root.Right, nil)
	}
	return root
}

// 递归2
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectHelper(root.Left, root.Right)
	return root
}

func connectHelper(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectHelper(node1.Left, node1.Right)
	connectHelper(node1.Right, node2.Left)
	connectHelper(node2.Left, node2.Right)
}

// 非递归 - 迭代
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		p := make([]*Node, 0) // 记录下一层节点
		for i, node := range q {
			if i < len(q)-1 {
				node.Next = q[i+1]
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p // 更新q为下一层节点
	}
	return root
}
