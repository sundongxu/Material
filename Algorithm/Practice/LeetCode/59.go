package LeetCode

/**
 * 59. Spiral Matrix II
 * 描述：
 * 难度：Medium
 * 类型：Matrix
 */
func GenerateMatrix(n int) [][]int {
	left, up := 0, 0
	right, down := n-1, n-1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	cur := 1
	for true {
		// 先从左往右遍历
		for i := left; i <= right; i++ {
			res[up][i] = cur
			cur++
		}
		up++
		if up > down {
			break
		}
		// 再从上往下遍历
		for i := up; i <= down; i++ {
			res[i][right] = cur
			cur++
		}
		right--
		if right < left {
			break
		}
		// 再从右往左遍历
		for i := right; i >= left; i-- {
			res[down][i] = cur
			cur++
		}
		down--
		if down < up {
			break
		}
		// 最后从下往上遍历
		for i := down; i >= up; i-- {
			res[i][left] = cur
			cur++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
