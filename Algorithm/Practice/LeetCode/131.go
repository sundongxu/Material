package LeetCode

/**
 * 131. Palindrome Partitioning
 * 描述：
 * 难度：Medium
 * 类型：String & Backtracking
 */
func partition(s string) [][]string {
	result, path := make([][]string, 0), make([]string, 0)
	dfs(s, path, &result, 0, 1)
	return result
}

// prev表示前一个隔板，start表示当前隔板
func dfs(s string, path []string, result *[][]string, prev, cur int) {
	if cur == len(s) {
		// 终止条件
		if isPalindrome(s, prev, cur-1) {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			pathCopy = append(pathCopy, s[prev:cur])
			*result = append(*result, pathCopy)
		}
		return
	}
	// 不断开
	// 如果s[prev:cur]是回文，则start处可以断开，并由此推进prev=cur，cur=cur+1，也可以不断开（上一行已经做了）
	dfs(s, path, result, prev, cur+1)
	if isPalindrome(s, prev, cur-1) {
		// 断开
		path = append(path, s[prev:cur])
		dfs(s, path, result, cur, cur+1)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end && s[start] == s[end] {
		start++
		end--
	}
	return start >= end
}
