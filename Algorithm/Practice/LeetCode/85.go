package LeetCode

/**
 * 85. Maximal Rectangle
 * 描述：
 * 难度：Hard
 * 类型：Matrix
 */
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	h, l, r := make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		r[i] = n
	}

	result := 0
	for i := 0; i < m; i++ {
		left, right := 0, n
		// 从左到右，计算l(i,j)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				h[j]++
				l[j] = max(l[j], left)
			} else {
				left = j + 1
				h[j], l[j], r[j] = 0, 0, n
			}
		}
		// 从右到左，计算r(i,j)
		for k := n - 1; k >= 0; k-- {
			if matrix[i][k] == '1' {
				r[k] = min(r[k], right)
				result = max(result, h[k]*(r[k]-l[k]))
			} else {
				right = k
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
