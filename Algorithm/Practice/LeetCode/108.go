package LeetCode

/**
 * 108. Convert Sorted Array to Binary Search Tree
 * 描述：
 * 难度：Medium
 * 类型：Tree & Recursion & DFS
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(0, len(nums), nums)
}

func sortedArrayToBSTHelper(begin, end int, nums []int) *TreeNode {
	distance := end - begin
	if distance <= 0 { // left <= right
		return nil
	}

	mid := begin + distance/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTHelper(begin, mid, nums)
	root.Right = sortedArrayToBSTHelper(mid+1, end, nums)
	return root
}

// 法二，二分法的变形
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(0, len(nums)-1, nums)
}

func sortedArrayToBSTHelper(begin, end int, nums []int) *TreeNode {
	distance := end - begin
	if distance < 0 { // left < right
		return nil
	}

	mid := begin + distance/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTHelper(begin, mid-1, nums)
	root.Right = sortedArrayToBSTHelper(mid+1, end, nums)
	return root
}

// left < right，终止条件是left == right，搜索范围为[left, right)，没搜到right是因为left无法大于right（两者相等的时候就已经终止了），已经满足终止条件了，所以right侧是开区间
// left <= right，终止条件是left > right即left === right + 1，搜索范围为[left, right]，很明显left都超过了right，搜索范围在right侧是闭区间
