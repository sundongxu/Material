package LeetCode

/**
 * 87. Scramble String
 * 描述：
 * 难度：Hard
 * 类型：DP
 */
// 动规
// 令f(s1,s2)表示s1和s2是否和谐
// 情况1：对于l(s1)和r(s1)没有被交换的情况，那么s2同样需要被分为s2(0,i)以及s2(i,n-i)
// 此时状态转移方程为f(s1,s2) = f(s1(0,i),s2(0,i)) && f(s1(i,n-i),s2(i,n-i))，i从1到n-1，多个表达式做或操作为最终值
// // 情况2：对于l(s1)和r(s1)被交换的情况，那么s2需要被分为s2(0,n-i)以及s2(n-i,i)
// 此时状态转移方程为f(s1,s2) = f(s1(0,i),s2(n-i,i)) && f(s1(i,n-i),s2(0,n-i))，i从1到n-1，多个表达式做或操作为最终值
// 最终 f(s1,s2) = 情况1 || 情况2
func isScramble(s1, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	// 第一个字符串从 i1 开始，第二个字符串从 i2 开始，子串的长度为 length
	// 和谐返回 1，不和谐返回 0
	var dfs func(i1, i2, length int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()

		// 判断两个子串是否相等
		x, y := s1[i1:i1+length], s2[i2:i2+length]
		if x == y {
			return 1
		}

		// 判断是否存在字符 c 在两个子串中出现的次数不同
		freq := [26]int{}
		for i, ch := range x {
			freq[ch-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq[:] {
			if f != 0 {
				return 0
			}
		}

		// 枚举分割位置
		for i := 1; i < length; i++ {
			// 不交换的情况
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}
			// 交换的情况
			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}

		return 0
	}
	return dfs(0, 0, n) == 1
}
