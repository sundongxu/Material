package LeetCode

// 144. Binary Tree Preorder Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	preorderTraversalRecursion(root, &res)
	return res
}

func preorderTraversalRecursion(root *TreeNode, sequence *[]int) {
	if root == nil {
		return
	}
	*sequence = append(*sequence, root.Val)
	preorderTraversalRecursion(root.Left, sequence)
	preorderTraversalRecursion(root.Right, sequence)
}

// 非递归 - 栈
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

// 非递归 - Morris算法
