package LeetCode

// 139. Word Break
// 令f[i]表示s[0,i]是否可以分词，可分词为true，否则为false
// 状态转移方程为：f[i] = any_of(f[j] && s[j+1, i] => dict)，0<=j<i
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && findWord(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func findWord(word string, wordDict []string) bool {
	for _, w := range wordDict {
		if w == word {
			return true
		}
	}
	return false
}

//// dfs
//func wordBreak(s string, wordDict []string) bool {
//	return dfs(&s, &wordDict, 0, 0)
//}
//
//func dfs(s *string, wordDict *[]string, start int, cur int) bool{
//	if cur == len(*s) {
//
//	}
//}
