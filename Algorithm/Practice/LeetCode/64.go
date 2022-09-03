package LeetCode

/**
 * 64. Minimum Path Sum
 * 描述：
 * 难度：Medium
 * 类型：DP
 */
// 令f[i][j]为走到点(i,j)的最小路径和
// 状态转移方程为：f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	// 计算第一行
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	// 计算第一列
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
