package LeetCode

// 44. Wildcard Matching
// 本题中'?'与正则表达式'.'作用相同，都用于匹配任意单个字符
// 本题中'*'则与正则表达式'*'的作用不同
// 后者只能让前一个字符出现0或多次，但是前者可以不依赖于前一个字符而匹配任意字符串
// 如'abcdefg'这样的任意字符串，在本题中都可以仅仅通过一个'*'来进行匹配
func isMatch(s string, p string) bool {
	return isMatchHelper(s, p, 0, 0)
}

func isMatchHelper(s, p string, si, pi int) bool {
	star := false
	ssi, ppi := -1, -1
	for ssi, ppi = si, pi; ssi < len(s) && ppi < len(p); ssi, ppi = ssi+1, ppi+1 {
		switch p[ppi] {
		case '?':
			break
		case '*':
			star = true
			si = ssi
			pi = ppi
			for p[pi] == '*' {
				// 跳过连续*
				pi++
			}
			// pi指向'*'后一个元素
			if pi >= len(p) {
				// 模式已经到末尾，且此前至少有一个'*'，一定匹配成功
				return true
			}
			ssi = si - 1
			ppi = pi - 1
			break
		default:
			if s[ssi] != p[ppi] {
				if !star {
					return false
				}
				si++
				ssi = si - 1
				ppi = pi - 1
			}
		}
	}
	for ppi < len(p) && p[ppi] == '*' {
		ppi++
	}
	return ppi >= len(p)
}

// 动态规划
// 令dp[i][j]为s的前i个字符和p的前j个字符是否匹配
// dp[i][j]的语义是：s的前i个字符和p的前j个字符是否匹配。
// 边界条件dp[0][0]=true，dp[0][j]=true(p的前j个字符都是*）
// 我们可以求解dp[i][j]:
// (1) 如果p[j]是字母，则dp[i][j]=(p[j-1]==s[i-1]) && dp[i-1][j-1]
// (2) 如果p[j]是?，则p[j]匹配任意一个字符，dp[i][j]=dp[i-1][j-1]
// (3) 如果p[j]是*，分为两种情况：
// a. 如果*代表0个字符的话，则s的前i个字符和p的前j个字符是否匹配，取决于s的前i个字符和p的前j-1个字符是否匹配，即dp[i][j]=dp[i][j-1]
// b. 如果*代表多个字符的话，则s的前i个字符和p的前j个字符是否匹配，取决于s的前i-1个字符和p的前j个字符是否匹配。
//    如果dp[i-1][j]=true，则dp[i][j]=true，此时即*代表的字符序列增加了一个字符s[i]而已，而*正好可以可以代表多个不同字符组成的字符串
// 综上a、b两种情况，可得dp[i][j]=dp[i-1][j] || dp[i][j-1]
func isMatch(s string, p string) bool {
	lenS, lenP := len(s)+1, len(p)+1
	dp := make([][]bool, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]bool, lenP)
	}
	dp[0][0] = true
	for j := 1; j < lenP; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < lenS; i++ {
		for j := 1; j < lenP; j++ {
			switch p[j-1] {
			case '*':
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			case '?', s[i-1]:
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[lenS-1][lenP-1]
}
