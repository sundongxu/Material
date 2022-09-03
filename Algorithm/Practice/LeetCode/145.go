package LeetCode

/**
 * 145. Binary Tree Postorder Traversal
 * 描述：
 * 难度：Easy
 * 类型：Tree & Recursion & Stack
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	postorderTraversalRecursion(root, &res)
	return res
}

func postorderTraversalRecursion(root *TreeNode, sequence *[]int) {
	if root == nil {
		return
	}
	postorderTraversalRecursion(root.Left, sequence)
	postorderTraversalRecursion(root.Right, sequence)
	*sequence = append(*sequence, root.Val)
}

// 非递归 - 栈
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var p, q *TreeNode
	p, q = root, nil
	for len(stack) > 0 {
		for p != nil {
			// 往左下走
			stack = append(stack, p)
			p = p.Left
		}
		q = nil
		for len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if p.Right == q {
				res = append(res, p.Val)
				q = p // 保存刚刚访问过的节点
			} else {
				// 当前节点不能访问，需要二次进栈
				stack = append(stack, p)
				// 先处理右子树
				p = p.Right
				break
			}
		}
	}
	return res
}
