package LeetCode

// 时间复杂度：O(m*n)
// 空间复杂度：O(m+n)
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if row[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if col[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

// 时间复杂度：O(m*n)
// 空间复杂度：O(1)
func setZeroes2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rowHasZero, colHasZero := false, false

	// 遍历第一行，看是不是有0
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowHasZero = true
		}
	}

	// 遍历第一列，看是不是有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colHasZero = true
		}
	}

	// 遍历除了第一行、第一列以外的全部位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// 如果某个位置为0，则将其所在行、所在列的第一个位置都置为0，作为一个标记
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 遍历除了第一行、第一列以外的全部位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				// 如果某个位置所在行或所在列的第一个元素为0，说明是此前标记，此时应该将该位置也置为0
				matrix[i][j] = 0
			}
		}
	}

	// 上面没有处理第一行和第一列的位置
	// 根据rowHasZero决定是不是要把第一行全部置为0
	if rowHasZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if colHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
