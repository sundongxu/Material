package LeetCode

/**
 * 171. Excel Sheet Column Number
 * 描述：
 * 难度：Easy
 * 类型：Math
 */
// 168题的逆运算
func titleToNumber(s string) int {
	val, res := 0, 0
	for i := 0; i < len(s); i++ {
		val = int(s[i] - 'A' + 1)
		res = res*26 + val
	}
	return res
}
