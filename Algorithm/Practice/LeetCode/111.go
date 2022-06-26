package LeetCode

// 111. Minimum Depth of Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		// 左右孩子都没有，说明当前节点root已经是叶子节点，那么以它为根的子树最小深度为1
		return 1
	}
	if root.Left == nil {
		// 没有左孩子，但有右孩子，那么以root为根的子树深度最小为其右子树的最小深度+1
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		// 没有右孩子，但有左孩子，那么以root为根的子树深度最小为其左子树的最小深度+1
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
