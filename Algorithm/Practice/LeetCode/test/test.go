package main

import "fmt"

func main() {
	// [3,9,20,null,null,15,7]
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node1}
	node3 := &TreeNode{Val: 1, Right: node2}
	result := postorderTraversal(node3)
	fmt.Printf("result = %+v", result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归 - 栈
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var p, q *TreeNode
	p, q = root, nil
	for {
		for p != nil {
			// 往左下走
			stack = append(stack, p)
			p = p.Left
		}
		q = nil
		for len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if p.Right == q {
				res = append(res, p.Val)
				q = p // 保存刚刚访问过的节点
			} else {
				// 当前节点不能访问，需要二次进栈
				stack = append(stack, p)
				// 先处理右子树
				p = p.Right
				break
			}
		}
		if len(stack) == 0 {
			break
		}
	}
	return res
}
