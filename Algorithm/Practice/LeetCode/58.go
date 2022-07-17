package LeetCode

// 58. Length of Last Word
func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last--
	}
	// 此时last指向字符串末尾最后一个非空格字符，即最后一个单词的末尾字符
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}
	return last - first
}
