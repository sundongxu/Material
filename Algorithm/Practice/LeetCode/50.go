package LeetCode

// 50. Pow(x, n)
// x^n = x^(n/2) * x^(n/2) * x^(n%2)
// n%2=0(n为偶数)，1(n为奇数)
// 所以n为偶数时，x^n = x^(n/2) * x^(n/2)
// 所以n为奇数时，x^n = x^(n/2) * x^(n/2) * x
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / power(x, -n)
	}
	return power(x, n)
}

func power(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	v := power(x, n/2)
	if n%2 == 0 {
		return v * v
	} else {
		return v * v * x
	}
}
