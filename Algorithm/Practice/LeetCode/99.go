package LeetCode

// 99. Recover Binary Search Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历，很好理解，空间复杂度O(n)
// https://leetcode.cn/problems/recover-binary-search-tree/solution/hui-fu-er-cha-sou-suo-shu-by-leetcode-solution/
func recoverTree(root *TreeNode) {
	nums := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwapped(nums)
	recover(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	x, y := nums[index1], nums[index2]
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Right, count, x, y)
	recover(root.Left, count, x, y)
}

// 由于我们只关心中序遍历的值序列中每个相邻的位置的大小关系是否满足条件
// 且错误交换后最多两个位置不满足条件，因此在中序遍历的过程我们只需要维护当前中序遍历到的最后一个节点prev
// 然后在遍历到下一个节点的时候，看两个节点的值是否满足前者小于后者即可
// 如果不满足说明找到了一个交换的节点，且在找到两次以后就可以终止遍历
func recoverTree(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	var x, y, prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && root.Val < prev.Val {
			y = root
			if x == nil {
				x = prev
			} else {
				break
			}
		}
		prev = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

// Morris中序遍历，不需要辅助栈，空间复杂度O(1)
func recoverTree(root *TreeNode) {
	var x, y, pred, predecessor *TreeNode

	for root != nil {
		if root.Left != nil {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else { // 说明左子树已经访问完了，我们需要断开链接
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predecessor.Right = nil
				root = root.Right
			}
		} else { // 如果没有左孩子，则直接访问右孩子
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func recoverTree(root *TreeNode) {
	var prev, target1, target2 *TreeNode
	_, target1, target2 = inOrderTraverse(root, prev, target1, target2)
	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}

}

func inOrderTraverse(root, prev, target1, target2 *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return prev, target1, target2
	}
	// 先访问左子树
	prev, target1, target2 = inOrderTraverse(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			// 仅在首次发现错位点时，更新第一个位点target1
			target1 = prev
		}
		// 每次都更新第二个位点target2，这是因为可能发现至多2个错位点
		target2 = root
	}
	// 类似中序遍历时，访问根节点
	prev = root
	// 最后访问右子树
	prev, target1, target2 = inOrderTraverse(root.Right, prev, target1, target2)
	return prev, target1, target2
}
