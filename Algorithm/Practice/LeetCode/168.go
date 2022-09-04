package LeetCode

/**
 * 168. Excel Sheet Column Title
 * 描述：
 * 难度：Easy
 * 类型：Math
 */
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
//...
// 26进制
func convertToTitle(n int) string {
	result := make([]byte, 0)
	for n > 0 {
		result = append(result, 'A'+byte((n-1)%26))
		n = (n - 1) / 26
	}
	// 反转数组
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
