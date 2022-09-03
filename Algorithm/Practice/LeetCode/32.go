package LeetCode

/**
 * 32. Longest Valid Parentheses
 * 描述：
 * 难度：Hard
 * 类型：Stack
 */
// 利用栈
// 这里需要计算嵌套括号的总长度，所以栈里面不能单纯的存左括号，而应该存左括号在原字符串的下标，这样通过下标相减可以获取长度
func longestValidParentheses(s string) int {
	stack, res := make([]int, 0), 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func longestValidParentheses(s string) int {
	stack, res := make([]int, 0), 0 // 栈用于存储未匹配的左括号的下标
	last := -1                      // 上一个未匹配的右括号下标
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 遇到左括号，就将其下标压入栈中
			stack = append(stack, i)
		} else {
			// 遇到右括号
			if len(stack) == 0 {
				// 栈为空，说明此时没有左括号可以用于匹配了
				last = i // 更新未匹配的右括号下标
			} else {
				// 栈非空，说明此时还有左括号可以用于匹配
				stack = stack[:len(stack)-1] // 将栈顶元素即最近一个左括号的下标出栈
				if len(stack) == 0 {
					res = max(res, i-last)
				} else {
					res = max(res, i-stack[len(stack)-1])
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses1(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

// 双指针法
func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	// 从左往右遍历
	// 会漏掉一种情况：左括号的数量始终大于右括号的数量，即 (()，但考虑了右括号的数量始终大于左括号的情况
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++ // 左括号计数
		} else {
			right++ // 右括号计数
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	// 从右往左遍历
	// 会漏掉一种情况：右括号的数量始终大于左括号的数量，即 ())，但考虑了左括号的数量始终大于右括号的情况
	// 相当于对上一轮循环的补偿，两个循环综合就能得到最终结果
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
