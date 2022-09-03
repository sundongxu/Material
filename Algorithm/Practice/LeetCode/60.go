package LeetCode

import "strconv"

/**
 * 60. Permutation Sequence
 * 描述：
 * 难度：Hard
 * 类型：Math & Recursion
 */
// 康托展开式：X = a[n]*(n-1)! + a[n-1]*(n-2)! + … + a[i]*(i-1)! + … + a[2]*1! + a[1]*0!
func factorial(i int) int {
	f := 1
	for ; i >= 1; i-- {
		f *= i
	}
	return f
}

func getPermutationHelper(s string, k int) string {
	i := len(s)
	if 1 == i {
		return s
	}
	factorial := factorial(i)
	nextFactorial := factorial / i
	if k <= nextFactorial {
		return s[:1] + getPermutationHelper(s[1:], k)
	}
	c, k := (k-1)/nextFactorial, (k-1)%nextFactorial+1
	if c > 0 {
		s = string(s[c]) + s[:c] + s[c+1:]
	}
	return getPermutationHelper(s, k)
}

func getPermutation(n int, k int) string {
	s := ""
	for i := 1; i <= n; i++ {
		s += strconv.Itoa(i)
	}
	return getPermutationHelper(s, k)
}

// https://leetcode.cn/problems/permutation-sequence/solution/di-kge-pai-lie-by-leetcode-solution/
func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans
}
