package LeetCode

import "strings"

// 28. Implement strStr()
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		// i表示父串本轮起始下标
		for j := 0; ; j++ {
			// j表示相对于i的下标位移，也可以认为是子串的当前下标
			if j == len(needle) {
				// 找到了，返回父串下标
				return i
			}
			if i+j == len(haystack) {
				// 父串起始下标+位移超过了父串长度
				// 说明剩余父串的长度已经小于子串的长度，此时已经不可能找到子串了
				return -1
			}
			if haystack[i+j] != needle[j] {
				// 父串当前下标对应字符 和 子串当前下标对应字符 不相等
				// 说明父串以i起始的子串不匹配目标子串，跳出本轮循环，重新从父串的下一个位置即i+1开始匹配
				break
			}
		}
	}
}
