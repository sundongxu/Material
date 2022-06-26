package LeetCode

// 63. Unique Paths II

// 令f[i][j]为走到点(i,j)的路径条数
// f[0][j]代表第0行，第j列，因为只能一直往右走，所以能直接算出，如果中途有一个障碍，那么之后列都是0
// f[i][0]代表第0列，第i行，因为只能一直往下走，所以也能直接算出，如果中途有一个障碍，那么之后行都是0
// 状态转移方程为：
// f[i][j] = 0，当obstacleGrid[i][j] == 1
// f[i][j] = f[i-1][j] + f[i][j-1]，当obstacleGrid[i][j] == 0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 计算第一行
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			for j := i; j < n; j++ {
				dp[0][j] = 0
			}
			break
		}
		dp[0][i] = 1
	}
	// 计算第一列
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			for j := i; j < m; j++ {
				dp[i][0] = 0
			}
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
