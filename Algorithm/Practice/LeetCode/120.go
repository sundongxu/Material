package LeetCode

import "math"

/**
 * 120. Triangle
 * 描述：
 * 难度：Medium
 * 类型：DP
 */
// 令f[i][j]为到点(i,j)的最小路径和
// 状态转移方程为：f[i][j] = max(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][len(dp[i])-1] = dp[i-1][len(dp[i-1])-1] + triangle[i][len(dp[i])-1]
	}
	for i := 2; i < m; i++ {
		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}
	minPathSum := math.MaxInt
	for j := 0; j < len(dp[m-1]); j++ {
		if dp[m-1][j] < minPathSum {
			minPathSum = dp[m-1][j]
		}
	}
	return minPathSum
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
