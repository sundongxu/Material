package LeetCode

/**
 * 101. Symmetric Tree
 * 描述：
 * 难度：Easy
 * 类型：Tree & Recursion
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecursion(root.Left, root.Right)
}

func isSymmetricRecursion(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricRecursion(left.Left, right.Right) && isSymmetricRecursion(left.Right, right.Left)
}
