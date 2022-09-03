package LeetCode

/**
 * 54. Spiral Matrix
 * 描述：
 * 难度：Medium
 * 类型：Matrix
 */
func SpiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	left, up := 0, 0
	right, down := len(matrix[0])-1, len(matrix)-1
	for true {
		// 先从左往右遍历
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		// 再从上往下遍历
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 再从右往左遍历
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		// 最后从下往上遍历
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
