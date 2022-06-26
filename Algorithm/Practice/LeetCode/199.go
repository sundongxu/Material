package LeetCode

// 199. Binary Tree Right Side View
// 本质是层次遍历取每层最后一个元素
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[l-1].Val)
		queue = queue[l:]
	}
	return res
}
