package LeetCode

import (
	"strings"
)

/**
 * 140. Word Break II
 * 描述：
 * 难度：Hard
 * 类型：DP
 */
// 同139，令f[i]表示s[0,i]是否可以分词，可分词为true，否则为false
// 状态转移方程为：f[i] = any_of(f[j] && s[j+1, i] => dict)，0<=j<i
// 另令prev[i][j]表示s[j,i)是否是一个合法单词，合法则为true，否则为false，可以从j处切开
func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	prev := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		prev[i] = make([]bool, n)
	}
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && findWord(s[j:i], wordDict) {
				dp[i] = true
				prev[i][j] = true
			}
		}
	}
	result, path := make([]string, 0), make([]string, 0)
	dfs(s, prev, len(s), path, &result)
	return result
}

func findWord(word string, wordDict []string) bool {
	for _, w := range wordDict {
		if w == word {
			return true
		}
	}
	return false
}

func dfs(s string, prev [][]bool, cur int, path []string, result *[]string) {
	if cur == 0 {
		tmp := strings.Join(reverse(path), " ")
		*result = append(*result, tmp)
	}
	for i := 0; i < len(s); i++ {
		if prev[cur][i] {
			path = append(path, s[i:cur])
			dfs(s, prev, i, path, result)
			path = path[0 : len(path)-1]
		}
	}
}

func reverse(strs []string) []string {
	// 建议是不要改变源数组，返回一个新的，因为path是还会后续复用的
	newStrs := make([]string, len(strs))
	copy(newStrs, strs)
	for i, j := 0, len(newStrs)-1; i < j; {
		newStrs[i], newStrs[j] = newStrs[j], newStrs[i]
		i++
		j--
	}
	return newStrs
}
