package LeetCode

/**
 * 202. Happy Number
 * 描述：
 * 难度：Easy
 * 类型：Math
 */
func isHappy(n int) bool {
	seq := make([]int, 0)
	for n != 1 {
		seq = append(seq, n)
		n = getSquareOfDigits(n)
		for _, num := range seq {
			if n == num {
				return false
			}
		}
	}
	return true
}

func getSquareOfDigits(n int) int {
	result := 0
	tmp := n
	for tmp != 0 {
		remainder := tmp % 10
		result += remainder * remainder
		tmp /= 10
	}
	return result
}
