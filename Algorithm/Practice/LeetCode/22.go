package LeetCode

/**
 * 22. Generate Parentheses
 * 描述：
 * 难度：Medium
 * 类型：DFS
 */
// 【DFS】
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]string, 0)
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lNum, rNum int, str string, res *[]string) {
	if lNum == 0 && rNum == 0 {
		*res = append(*res, str)
		return
	}
	if lNum > 0 {
		findGenerateParenthesis(lNum-1, rNum, str+"(", res)
	}
	if rNum > 0 && rNum > lNum {
		findGenerateParenthesis(lNum, rNum-1, str+")", res)
	}
}
