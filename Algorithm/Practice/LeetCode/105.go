package LeetCode

// 105. Construct Binary Tree from Preorder and Inorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 找根节点，一定是前序遍历序列的第一个元素
	root := &TreeNode{Val: preorder[0]}
	// 找到以root为根的树，其左子树的前序、中序遍历序列，及其右子树的前序、中序遍历序列
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			// 在中序遍历序列中找到根节点下标
			// 分别递归构造root的左子树和右子树
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return root
}
