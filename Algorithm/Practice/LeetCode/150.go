package LeetCode

import "strconv"

/**
 * 150. Evaluate Reverse Polish Notation
 * 描述：
 * 难度：Medium
 * 类型：
 */
// 逆波兰表达式严格遵循「从左到右」的运算
// 计算逆波兰表达式的值时，使用一个栈存储操作数，从左到右遍历逆波兰表达式，进行如下操作：
//（1）如果遇到操作数，则将操作数入栈
//（2）如果遇到运算符，则将两个操作数出栈，其中先出栈的是右操作数，后出栈的是左操作数，使用运算符对两个操作数进行运算，将运算得到的新操作数入栈
// 整个逆波兰表达式遍历完毕之后，栈内只有一个元素，该元素即为逆波兰表达式的值
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			// 当前token是数字，则入栈
			stack = append(stack, val)
		} else {
			// 当前token是运算符，则出栈两个元素作为运算符，且先出栈的是右操作数，后出栈的是左操作数，运算结果重新压栈
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
