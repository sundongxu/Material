package LeetCode

/**
 * 103. Binary Tree Zigzag Level Order Traversal
 * 描述：
 * 难度：Medium
 * 类型：Tree & BFS & Recursion
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归 - DFS
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	zigzagLevelOrderRecursion(root, 1, true, &res)
	return res
}

func zigzagLevelOrderRecursion(root *TreeNode, level int, left2Right bool, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) < level {
		*res = append(*res, []int{})
	}

	(*res)[level-1] = append((*res)[level-1], root.Val)
	// 如果是从右往左，那么把最后一个元素提前到第一个位置，相当于从头部插入
	if !left2Right {
		n := len((*res)[level-1])
		copy((*res)[level-1][1:n], (*res)[level-1][0:n-1])
		(*res)[level-1][0] = root.Val
	}

	zigzagLevelOrderRecursion(root.Left, level+1, !left2Right, res)
	zigzagLevelOrderRecursion(root.Right, level+1, !left2Right, res)
}

// 非递归 - BFS
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0) // 队列
	queue = append(queue, root)
	left2Right := true
	for len(queue) != 0 {
		level := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			// 遍历此时队列全部元素，即一层元素
			front := queue[i]
			level = append(level, front.Val)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		// 将当前这一层的元素全部出队，此时队列仅剩下一层全部元素
		queue = queue[l:]
		// 反转本层结果
		if !left2Right {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		left2Right = !left2Right
		res = append(res, level)
	}
	return res
}
