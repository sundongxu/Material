package LeetCode

// 69. Sqrt(x)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right, lastMid := 1, x/2, 0
	for left <= right {
		mid := left + (right-left)/2
		if x/mid == mid {
			return mid
		} else if x/mid < mid {
			// 此时无需更新lastMid，因为mid的平方大于x，无论如何也不会把该mid作为x的平方根的最终结果
			right = mid - 1
		} else {
			// 此时虽然要将左边界移至中点mid以后，但是由于这里存在平方根可能不是整数的问题，所以可能会需要四舍五入
			// 所以要更新下lastMid，万一后续无法找到更加合适的平方根(平方比x小，但比lastMid的平方大，即比lastMid更接近于x的平方根)，那就返回它
			lastMid = mid
			left = mid + 1
		}
	}
	return lastMid
}
