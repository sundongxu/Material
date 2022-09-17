package LeetCode

/**
 * 129. Sum Root to Leaf Numbers
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}
	return dfs(root.Left, sum*10+root.Val) + dfs(root.Right, sum*10+root.Val)
}
