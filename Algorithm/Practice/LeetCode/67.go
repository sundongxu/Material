package LeetCode

import "strconv"

/**
 * 67. Add Binary
 * 描述：
 * 难度：Easy
 * 类型：Array & Math
 */
// 难点在于如何去除前导0，因为可能有溢出问题，所以不能直接转为Atoi再Itoa
func addBinary(a string, b string) string {
	totalLen := max(len(a), len(b)) + 1
	arr := make([]int, totalLen)
	for i := 0; i < totalLen-1; i++ {
		ai, bi := len(a)-i-1, len(b)-i-1
		sum := 0
		if ai >= 0 && bi >= 0 {
			// a、b都未到起始字符
			sum = arr[totalLen-i-1] + int(a[ai]-'0') + int(b[bi]-'0')

		} else if ai >= 0 {
			// a长度大于b
			sum = arr[totalLen-1-i] + int(a[ai]-'0')
		} else {
			// b长度大于a
			sum = arr[totalLen-1-i] + int(b[bi]-'0')
		}
		if sum >= 2 {
			arr[totalLen-i-2]++
		}
		arr[totalLen-i-1] = sum % 2
	}
	var result string
	for _, bit := range arr {
		result += strconv.Itoa(bit)
	}
	i, count := 0, 0
	for i < len(result) {
		if result[i] == '0' {
			count++
			i++
		} else {
			break
		}
	}
	if count == len(result) {
		return "0"
	}
	return result[count:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
