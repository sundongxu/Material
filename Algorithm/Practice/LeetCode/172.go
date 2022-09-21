package LeetCode

/**
 * 172. Factorial Trailing Zeroes
 * 描述：
 * 难度：Medium
 * 类型：Math
 */
// n!尾零的数量即为n!中因子10的个数，而10=2×5，因此转换成求n!中质因子2的个数和质因子5的个数的较小值
// 由于质因子5的个数不会大于质因子2的个数，我们可以仅考虑质因子5的个数
// 而n!中质因子5的个数等于[1,n]的每个数的质因子5的个数之和，可以通过遍历[1,n]的所有5的倍数求出
func trailingZeroes(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}

func trailingZeroes(n int) (ans int) {
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return
}
