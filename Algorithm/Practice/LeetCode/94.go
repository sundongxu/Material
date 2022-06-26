package LeetCode

// 94. Binary Tree Inorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	inorderTraversalRecursion(root, &res)
	return res
}

func inorderTraversalRecursion(root *TreeNode, sequence *[]int) {
	if root == nil {
		return
	}
	inorderTraversalRecursion(root.Left, sequence)
	*sequence = append(*sequence, root.Val)
	inorderTraversalRecursion(root.Right, sequence)
}

// 非递归 - 栈 -> slice
// 入栈: stack = append(stack, node)，插入末尾
// 出栈: stack = stack[:len(stack)-1]，从末尾弹出
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			// 如果当前节点不为空，则一直将其压栈
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// 非递归 - Morris算法
