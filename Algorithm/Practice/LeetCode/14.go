package LeetCode

/**
 * 14. Longest Common Prefix
 * 描述：
 * 难度：Easy
 * 类型：String
 */
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str0 := strs[0]
	longestPrefix := ""
	for i := range str0 {
		for _, str := range strs {
			if len(str) <= i || str[i] != str0[i] {
				return longestPrefix
			}
		}
		longestPrefix += string(str0[i])
	}
	return longestPrefix
}
