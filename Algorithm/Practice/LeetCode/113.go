package LeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, path := make([][]int, 0), make([]int, 0)
	pathSumRecursion(root, targetSum, path, &res)
	return res
}

func pathSumRecursion(root *TreeNode, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		pathCopy = append(pathCopy, root.Val)
		*res = append(*res, pathCopy)
		return
	}
	path = append(path, root.Val)
	if root.Left != nil {
		pathSumRecursion(root.Left, targetSum-root.Val, path, res)
	}
	if root.Right != nil {
		pathSumRecursion(root.Right, targetSum-root.Val, path, res)
	}
	path = path[:len(path)-1]
}
