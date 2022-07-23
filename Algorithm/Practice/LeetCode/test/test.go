package main

import "fmt"

func main() {
	// "aa"
	//"a"
	//"aa"
	//"a*"
	//"ab"
	//".*"
	//"ab"
	//".*c"
	//"a"
	//"ab*"
	// "a"
	// ".*..a*"
	//result := isMatch("aa", "a")
	//fmt.Println(result)
	//
	//result = isMatch("aa", "a*")
	//fmt.Println(result)
	//
	//result = isMatch("ab", ".*")
	//fmt.Println(result)
	//
	//result := isMatch("ab", ".*c")
	//fmt.Println(result)
	//
	//result = isMatch("a", "ab")
	//fmt.Println(result)

	result := isMatch("a", ".*..a*")
	fmt.Println(result)
}

func isMatch(s string, p string) bool {
	return isMatchHelper(s, p, 0, 0)
}

func isMatchHelper(s, p string, si, pi int) bool {
	if pi >= len(p) {
		// 模式匹配完成，看源字符串是否匹配完成
		return si >= len(s)
	}
	if pi < len(p)-1 && p[pi+1] != '*' {
		// 模式还剩至少一个字符
		// 模式下一个字符串不是'*'，则比较二者当前字符
		if si < len(s) && (p[pi] == s[si] || p[pi] == '.') {
			// 要么二者当前字符相等
			// 要么模式当前字符为'.'(匹配任意字符)且源字符串还未匹配完成
			// 源字符串与模式各往后推进一位
			return isMatchHelper(s, p, si+1, pi+1)
		} else {
			// 否则二者不匹配
			return false
		}
	} else if pi == len(p)-1 {
		// 已经匹配到模式的最后一个字符: pi == len(p)-1
		if si == len(s)-1 && (p[pi] == s[si] || p[pi] == '.') {
			// 源字符串也已经匹配到最后一个字符 且 (二者当前字符相等 或 模式最后一个字符是'.'
			return true
		}
		return false
	} else {
		// 模式还剩至少一个字符
		// 模式下一个字符为'*'，可能可以一直匹配源字符串的连续相同字符(1个或多个)，也可能直接跳过模式的两个字符(.* => .任意字符出现0次 => 空串)
		for si < len(s) && (p[pi] == s[si] || p[pi] == '.') {
			if isMatchHelper(s, p, si, pi+2) {
				// 因为模式下一个字符是'*'，那么其当前字符可以出现0次，模式字符串索引直接往后跳2位
				return true
			}
			// 模式当前字符一定要出现至少1次，可能是正好等于源字符串当前字符，也可以是'.'通配任何字符
			// 也可能源字符串的当前字符连续出现多次，那么可以一直往后推进源字符串，直到不匹配为止
			si++
		}
		// 模式当前字符既不与源字符串当前字符相等，也不是'.'通配符，那么只能让模式当前字符出现0次，即pi直接往后跳两位再接着与源字符串比较
		return isMatchHelper(s, p, si, pi+2)
	}
}
