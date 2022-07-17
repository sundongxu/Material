package LeetCode

import "math"

// 29. Divide Two Integers
func divide(dividend int, divisor int) int {
	var a, b, result int64
	a, b = int64(dividend), int64(divisor)
	if dividend < 0 {
		a = -a
	}
	if divisor < 0 {
		b = -b
	}
	for a >= b {
		c := b
		for i := 0; a >= c; i++ {
			a -= c
			result += 1 << i
			c <<= 1 // c = 2*c
		}
	}
	if (dividend^divisor)>>31 < 0 {
		return -int(result)
	}
	// dividend = math.MinInt32，divisor = -1 => quotient = math.MaxInt32 + 1 会溢出
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	return int(result)
}
