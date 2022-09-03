package LeetCode

/**
 * 132. Palindrome Partitioning II
 * 描述：
 * 难度：Hard
 * 类型：DP
 */
// 令f[i]为区间[i,n-1]之间最小的cut次数，n为字符串长度
// 状态转移方程为: f[i] = min{f(j+1)+1}, i<=j<=n
// 令p[i]为[j]区间[i,j]对应子串是否为回文串
// 状态转移方程为：p[i][j] = str[i] == str[j] && (j-i<2 || p[i+1][j-1])
func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	p := make([][]bool, n)
	for i := 0; i < n; i++ {
		p[i] = make([]bool, n)
	}
	for i := 0; i <= n; i++ {
		// 最差情况是每个字符都切一下
		// dp[0] = n-1
		dp[i] = n - 1 - i
	}
	// 再来看看有没有更好的情况
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || p[i+1][j-1]) {
				p[i][j] = true
				dp[i] = min(dp[i], dp[j+1]+1)
			}
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
