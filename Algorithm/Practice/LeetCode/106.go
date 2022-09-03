package LeetCode

/**
 * 106. Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 找根节点，一定是后序遍历序列的最后一个元素
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	// 找到以root为根的树，其左子树的前序、中序遍历序列，及其右子树的前序、中序遍历序列
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			// 在中序遍历序列中找到根节点下标
			// 分别递归构造root的左子树和右子树
			root.Left = buildTree(inorder[:i], postorder[:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}
	return root
}
