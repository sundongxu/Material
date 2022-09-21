package LeetCode

/**
 * 84. Largest Rectangle in Histogram
 * 描述：
 * 难度：Hard
 * 类型：Stack & MonoStack
 */

func largestRectangleArea(heights []int) int {
	n := len(heights)
	// left[i]代表以位置i的高度为基准，面积最大的矩形的最左端点的下标
	// right[i]代表以位置i的高度为基准，面积最大的矩形的最右端点的下标
	left, right := make([]int, n), make([]int, n)
	monoStack := make([]int, 0) // 元素为数组下标
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			// 栈顶元素的高度大于当前元素高度，则将栈顶元素出栈
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			// 栈为空，说明当前元素左边没有比它的高度小的元素了，那么以它的高度为基准，面积最大的矩形的左端点可以到-1(哨兵，类似链表的虚拟头指针dummy)
			left[i] = -1
		} else {
			// 栈不为空，说明当前元素左边有比它高度小的元素，且在它左边、离它最近的、高度比它小的元素，正是栈顶元素
			// 那么此时栈顶元素的下标，即以当前元素的高度为基准，面积最大的矩形的左端点
			left[i] = monoStack[len(monoStack)-1]
		}
		// 将当前元素入栈，即每次前面都把高度>=当前元素的高度的元素给出栈了，才来入栈
		// 说明该栈是一个元素依次递增(数组下标总是随着遍历而递增)，且元素高度也依次递增的栈，这种栈称为单调栈
		monoStack = append(monoStack, i)
	}
	monoStack = make([]int, 0) // 将栈重新初始化，用于下面计算right[]数组
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			// 栈为空，说明当前元素右边没有比它的高度小的元素了，那么以它的高度为基准，面积最大的矩形的右端点可以到n(哨兵)
			right[i] = n
		} else {
			// 栈不为空，说明当前元素右边有比它高度小的元素，且在它右边、离它最近的、高度比它小的元素，正是栈顶元素
			// 那么此时栈顶元素的下标，即以当前元素的高度为基准，面积最大的矩形的右端点
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
