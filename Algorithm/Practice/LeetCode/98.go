package LeetCode

import "math"

/**
 * 98. Validate Binary Search Tree
 * 描述：
 * 难度：Medium
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
// [5,4,6,null,null,3,7]
func isValidBST(root *TreeNode) bool {
	return isValidBSTRecursion(root, math.MinInt, math.MaxInt)
}

func isValidBSTRecursion(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	if root.Left != nil && (root.Val <= root.Left.Val) {
		return false
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}
	return isValidBSTRecursion(root.Left, min, root.Val) && isValidBSTRecursion(root.Right, root.Val, max)
}

// 根据二叉搜索树中序遍历序列一定是从小到大有序的这一性质
func isValidBST(root *TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
