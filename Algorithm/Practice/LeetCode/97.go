package LeetCode

/**
 * 97. Interleaving String
 * 描述：
 * 难度：Medium
 * 类型：DP
 */
// 动规
// 令f(i,j)表示s1的前i个元素和s2的前j个元素是否能交错构成s3的前i+j个元素
// 如果s1的第i个元素和s3的第i+j个元素相等
// 那么s1的前i个元素和s2的前j个元素是否能交错构成s3的前i+j个元素取决于
// s1的前i-1个元素和s2的前j个元素能否交错构成s3的前i+j-1个元素
// 即此时f(i,j)取决于f(i-1,j)，后者为true则前者也为true
// 同理，如果s2的第j个元素和s3的第i+j个元素相等并且f(i,j-1)为true，则f(i,j)也为true
// 于是可推导出状态转移方程：f(i,j) = [f(i-1,j) && s1(i-1) == s3(i+j-1)] || [f(i,j-1) && s2(j-1) == s3(i+j-1)]
// 其中边界条件为f(0,0)=true
// f(i,0) = s1(i-1) == s3(i-1) && f(i-1,0)，说明s2为空串，s1要整个等于s3才行
// f(0,j) = s2(j-1) == s3(j-1) && f(0,j-1)，说明s1为空串，s2要整个等于s3才行
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[n][m]
}
