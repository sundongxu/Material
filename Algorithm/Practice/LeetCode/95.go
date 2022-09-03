package LeetCode

/**
 * 95. Unique Binary Search Trees II
 * 描述：
 * 难度：Medium
 * 类型：Tree + Recursion
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesRecursion(1, n)
}

func generateTreesRecursion(start, end int) []*TreeNode {
	tree := make([]*TreeNode, 0)
	if start < end {
		tree = append(tree, nil)
		return tree
	}
	// start >= end
	for i := start; i <= end; i++ {
		left := generateTreesRecursion(start, i-1)
		right := generateTreesRecursion(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
