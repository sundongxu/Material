package LeetCode

import "strconv"

/**
 * 38. Count and Say
 * 描述：
 * 难度：Medium
 * 类型：String
 */
// 难度在于理解题意，本质是一个递推公式
// n = 1 => countAndSay(1) => 1，初始值
// n = 2 => countAndSay(2) => 如何读n=1的结果 => 如何读countAndSay(1) => 如何读1 => 1个1 => 11
// n = 3 => countAndSay(3) => 如何读n=2的结果 => 如何读countAndSay(2) => 如何读11 => 2个1 => 21
// n = 4 => countAndSay(4) => 如何读n=3的结果 => 如何读countAndSay(3) => 如何读21 => 1个2，1个1 => 1211
// n = 5 => countAndSay(5) => 如何读n=4的结果 => 如何读countAndSay(4) => 如何读1211 => 1个1，1个2，2个1 => 111221
// n = 6 => countAndSay(6) => 如何读n=5的结果 => 如何读countAndSay(5) => 如何读111221 => 3个1，2个2，1个1 => 312211
// ...
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return count(countAndSay(n - 1))
}

func count(s string) string {
	c := string(s[0]) // 从第一个字符开始计数
	count := 1        // 第一个字符已经出现一次，计数1
	res := ""
	for _, ch := range s[1:] {
		if string(ch) == c {
			// 仍然当前计数字符
			count += 1
		} else {
			// 已经不是当前计数字符，已经可以将当前计数字符的结果追加到res中，接下来再更新c和count
			res = res + strconv.Itoa(count) + c
			c = string(ch)
			count = 1
		}
	}
	res = res + strconv.Itoa(count) + c
	return res
}
