package LeetCode

/**
 * 48. Rotate Image
 * 描述：
 * 难度：Medium
 * 类型：Matrix
 */
// 方法一：先沿着副对角线翻转一次，再沿着水平中线翻转一次
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	// 沿着副对角线翻转
	for i := 0; i < n; i++ {
		// 遍历所有行
		for j := 0; j < n-i; j++ {
			// 遍历副对角线的左边所有列
			// (i,j)离左下角点的纵向距离为n-1-i，离左下角点的横向距离为j
			// 那么(i,j)关于副对角线对称的点，其离左下角点的横向距离应该为n-1-i，纵向距离应该为n-1-j，故而其坐标为(n-1-j, n-1-i)
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
	// 沿着水平中线翻转
	for i := 0; i < n/2; i++ {
		// 遍历一半的行
		for j := 0; j < n; j++ {
			// 遍历所有列
			// (i,j)关于水平中线对称的点坐标为(n-1-i,j)
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
}

// 方法二：先沿着水平中线翻转一次，再沿着主对角线翻转一次
func rotate2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	// 沿着水平中线翻转
	for i := 0; i < n/2; i++ {
		// 遍历一半的行
		for j := 0; j < n; j++ {
			// 遍历所有列
			// (i,j)关于水平中线对称的点坐标为(n-1-i,j)
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	// 沿着主对角线翻转
	for i := 0; i < n; i++ {
		// 遍历所有行
		for j := 0; j < i; j++ {
			// 遍历主对角线的左边所有列
			// (i,j)离左上角点的纵向距离是i，横向距离为j
			// 那么(i,j)关于主对角线对称的点，其离左上角点的横向距离应该为i，纵向距离应该为j，故而其坐标为(j, i)
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
