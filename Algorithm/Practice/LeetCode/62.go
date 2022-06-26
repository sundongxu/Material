package LeetCode

// 62. Unique Paths
// 动态规划
// 令f[i][j]为走到点(i,j)的路径条数
// f[0][j]代表第0行，第j列，因为只能一直往右走，所以能直接算出
// f[i][0]代表第0列，第i列，因为只能一直往下走，所以也能直接算出
// 状态转移方程为：f[i][j] = f[i-1][j] + f[i][j-1]，i, j > 0
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 计算第一行
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 计算第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
