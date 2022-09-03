package LeetCode

/**
 * 115. Distinct Subsequences
 * 描述：
 * 难度：Hard
 * 类型：DP
 */
// 动规
// 假设m和n分别为字符串s和t的长度
// 当m≥n时，可以通过动态规划的方法计算在s的子序列中t出现的个数
// 令f[i][j] 表示在s[i:]的子序列中t[j:]出现的个数
// 先考虑边界情况
// 当j=n时，t[j:]为空字符串，由于空字符串是任何字符串的子序列，因此f[i][j]=f[0..m][n]=1
// 当i=m且j<n时，s[i:]为空字符串，t[j:]为非空字符串，由于任何非空字符串都不是空字符串的子序列，因此f[i][j]=f[m][0..n-1]=0
// 再考虑非边界情况：i<m 且 j<n
// 当s[i]=t[j]时，f[i][j]=f[i+1][j+1](让s[i]和t[j]匹配)+f[i+1][j](不让s[i]和t[j]匹配，而让s[i]以后的字符与t[j]匹配)
// 当s[i]!=t[j]时，f[i][j]=f[i+1][j]
func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				//
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
