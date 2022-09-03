package LeetCode

/**
 * 107. Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	// 复用层次遍历算法，得到的结果res是自顶向下的
	res := levelOrder(root)
	// 反转最终结果res数组即可得到自底向上的结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// 递归 - DFS，难理解
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	levelOrderRecursion(root, 1, &res) // 起始层数为第一层
	return res
}

func levelOrderRecursion(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) < level {
		// 说明到了新的一层，新增容纳新层元素的数组
		*res = append(*res, []int{})
	}

	// 添加当前根节点值到其所在层对应的数组中
	(*res)[level-1] = append((*res)[level-1], root.Val)
	// 深度优先遍历，会一直往左走，虽然是会一直访问到左孩子，但是由于每次都会递增层数，所以每次都会把左孩子放进其所在层数对应的数组中
	levelOrderRecursion(root.Left, level+1, res)
	// 同上理
	levelOrderRecursion(root.Right, level+1, res)
	// 不考虑层次，上面的过程本质是：根 -> 左 -> 右，即二叉树的先序遍历，通过层级这一概念，将层次不同的元素放到不同数组，层次相同的元素追加到同一数组，就实现了层次遍历，神奇！
}
