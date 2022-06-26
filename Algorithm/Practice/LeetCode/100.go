package LeetCode

// 100. Same Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeRecursion(p, q)
}

func isSameTreeRecursion(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTreeRecursion(p.Left, q.Left) && isSameTreeRecursion(p.Right, q.Right)
}
