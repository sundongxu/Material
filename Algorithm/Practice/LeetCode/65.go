package LeetCode

/**
 * 65. Valid Number
 * 描述：
 * 难度：Hard
 * 类型：String & Math
 */
func isNumber(s string) bool {
	// 用一组标志位分别记录当前遍历字符前序字符的情况
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			// 当前字符为小数点，合法条件是之前没有小数点也没有e/E
			// 前序字符不能有e/E是因为题目规定了e/E之后必须是一个整数
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			// 当前字符为e/E，合法条件是之前没有e/E，且有数字位
			eFlag = true
			// 重置数位，因为e/E之后必须也要跟一个整数，否则仍不合法
			numFlag = false
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			// 符号位必须出现在第一个字符或前一个字符为e/E，才合法
			continue
		} else {
			return false
		}
	}
	// avoid case: s == '.' or 'e/E' or '+/-' and etc...
	// string s must have num
	return numFlag
}
