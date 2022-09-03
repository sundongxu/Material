package LeetCode

/**
 * 102. Binary Tree Level Order Traversal
 * 描述：
 * 难度：Medium
 * 类型：Tree & Queue & BFS
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Queue struct {
	nodes []*TreeNode
	num   int
}

func NewQueue() *Queue {
	return &Queue{
		nodes: make([]*TreeNode, 0),
		num:   0,
	}
}

func (q *Queue) enqueue(node *TreeNode) {
	q.nodes = append(q.nodes, node)
	q.num++
}

func (q *Queue) dequeue() *TreeNode {
	if q.num == 0 {
		return nil
	}
	front := q.nodes[0]
	q.nodes = q.nodes[1:]
	q.num--
	return front
}

func (q *Queue) len() int {
	return q.num
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := NewQueue()
	queue.enqueue(root)
	for queue.len() != 0 {
		level := make([]int, 0)
		num := queue.len()
		for i := 0; i < num; i++ {
			front := queue.dequeue()
			level = append(level, front.Val)
			if front.Left != nil {
				queue.enqueue(front.Left)
			}
			if front.Right != nil {
				queue.enqueue(front.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// 非递归 - BFS，好理解
// 用一个队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0) // 队列
	queue = append(queue, root)
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
		res = append(res, level)
	}
	return res
}

// 非递归 - BFS，好理解
// 用两个队列，一个存当前层的节点，一个存下一层的节点，遍历完当前层后，二者互换
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	curQueue, nextQueue := make([]*TreeNode, 0), make([]*TreeNode, 0)
	curQueue = append(curQueue, root)
	for len(curQueue) != 0 {
		level := make([]int, 0)
		l := len(curQueue)
		for i := 0; i < l; i++ {
			// 遍历此时队列全部元素，即一层元素
			front := curQueue[i]
			level = append(level, front.Val)
			if front.Left != nil {
				nextQueue = append(nextQueue, front.Left)
			}
			if front.Right != nil {
				nextQueue = append(nextQueue, front.Right)
			}
		}
		// 将下一层队列元素转交给当前层队列
		// 而curQ在上述循环后实际已经清空，所以交换二者后，实际还把nextQ也清空了
		curQueue, nextQueue = nextQueue, curQueue
		res = append(res, level)
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
