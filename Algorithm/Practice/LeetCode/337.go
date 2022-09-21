package LeetCode

/**
 * 337. House Robber III
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
// 房屋构成二叉树
// 一棵二叉树，树上的每个点都有对应的权值，每个点有两种状态（选中和不选中），问在不能同时选中有父子关系的点的情况下，能选中的点的最大权值和是多少
// 令f(i)表示选择i节点的情况下，i节点的子树上被选择的节点的最大权值和，不包括i节点的权值
// 令g(i)表示不选择i节点的情况下，i节点的子树上被选择的节点的最大权值和
// 令l和r代表i的左右孩子
// 当i被选中时，i的左右孩子都不能被选中，故i被选中情况下，子树上被选中点的最大权值和为l和r不被选中的最大权值和相加，即 f(i) = g(l) + g(r)
// 当i不被选中时，i的左右孩子可以被选中，也可以不被选中，也可以一个选中一个不选中
// 对于i的某个具体的孩子x，它对i的贡献是x被选中(即f(x))和不被选中(即g(x))情况下权值和的较大值
// 由于i的左右孩子可能都选中，也可能都不选中，还可能一个选中一个不选中，因此左右孩子对g(i)的贡献是独立的，而非互斥的
// 故g(i) = max{f(l), g(l)} + max{f(r), g(r)}
func rob(root *TreeNode) int {
	val := dfs(root)
	// val[0]对应f(root)即选中root，val[1]对应g(root)即不选root
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	// 选中当前节点，那么无法选中其左右孩子
	selected := node.Val + l[1] + r[1]
	// 不选中当前节点，那么左右孩子均可选可不选
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
