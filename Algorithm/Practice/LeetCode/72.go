package LeetCode

/**
 * 72. Edit Distance
 * 描述：
 * 难度：Hard
 * 类型：DP
 */
// 动规
// 令f[i][j]表示A[0,i]和B[0,j]之间的最小编辑距离
// 设A[0,i]的形式是str1c，B[0,j]的形式是str2d
// 如果 c == d，则f[i][j]=f[i-1][j-1]
// 如果 c != d，则
//（1）如果将c替换成d，则f[i][j]=f[i-1][j-1]+1
//（2）如果在c后面添加一个d，则f[i][j]=f[i][j-1]+1
//（3）如果将c删除，则f[i][j]=f[i-1][j]+1
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min([]int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]}...)
			}
		}
	}
	return dp[m][n]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
