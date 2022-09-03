package LeetCode

import "math"

/**
 * 8. String to Integer (atoi)
 * 描述：
 * 难度：Medium
 * 类型：String & Math
 */
func myAtoi(s string) int {
	num, sign, n := 0, 1, len(s)
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	// 提前判断，防止全部是0，导致下面逻辑访问数组越界
	if i == n {
		return 0
	}
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// 不能直接用math.MaxInt和math.MinInt，可能是2的64次方，与机器位数有关
	// 可以直接用math.MaxInt32和math.MinInt32
	//maxInt, minInt := 1<<31-1, -(1 << 31)
	maxInt, minInt := math.MaxInt32, math.MinInt32
	for ; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		if num > maxInt/10 || (num == maxInt/10 && int(s[i]-'0') > maxInt%10) {
			if sign == -1 {
				return minInt
			} else {
				return maxInt
			}
		}
		num = num*10 + int(s[i]-'0')
	}
	return sign * num
}
