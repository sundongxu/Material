package LeetCode

// 114. Flatten Binary Tree to Linked List
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 先打平左右子树，以root左右孩子为根的子树，就分别变成了一个符合题设条件的链表，每个树节点只有右孩子，左孩子都为空
	flatten(root.Left)
	flatten(root.Right)

	// 再来处理以root为根的子树
	// 如果root没有左孩子，说明以root为根的子树已经符合要求，直接返回
	if root.Left == nil {
		return
	}

	// 走到这里，说明root有左孩子
	// 将root.Left插入到root和root.Right之间
	// 并且将root.Right接在root左子树的最右节点的右孩子位置
	p := root.Left
	for p.Right != nil {
		// 寻找左孩子的最右节点，后续会将root.Right作为其右孩子
		p = p.Right
	}
	// 将root.Right插入成为root左子树的最右节点的右孩子
	p.Right = root.Right
	// 将root.Left插入成为root的右孩子
	root.Right = root.Left
	// 将root.Left置空
	root.Left = nil
}

// 非递归
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		top.Left = nil
		if len(stack) > 0 {
			top.Right = stack[len(stack)-1]
		}
	}
}

// 非递归 + 先序
func flatten(root *TreeNode) {
	sequence := make([]int, 0)
	preorder(root, &sequence)
	cur := root
	for i := 1; i < len(sequence); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{Val: sequence[i], Left: nil, Right: nil}
		cur = cur.Right
	}
	return
}

func preorder(root *TreeNode, sequence *[]int) {
	if root == nil {
		return
	}
	*sequence = append(*sequence, root.Val)
	preorder(root.Left, sequence)
	preorder(root.Right, sequence)
}
