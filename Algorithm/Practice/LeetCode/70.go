package LeetCode

// 70. Climbing Stairs
// 令f[i]为i级台阶的走法数，由于每次要么走1步，要么走2步
// 那么状态转移方程为f[i] = f[i-1] + f[i-2]
// 且f[0]=1,f[1]=1,f[2]=2=f[0]+f[1]
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
