package LeetCode

/**
 * 20. Valid Parentheses
 * 描述：
 * 难度：Easy
 * 类型：Stack
 */
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		switch c {
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				break
			} else {
				return false
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				break
			} else {
				return false
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				break
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
