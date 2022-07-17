package LeetCode

// 52. N-Queens II
func totalNQueens(n int) int {
	positions := make([]int, n) // 第i个元素的值positions[i]表示，第i行第positions[i]列放置皇后
	for i, _ := range positions {
		positions[i] = -1
	}
	count := 0
	solveNQueensHelper(positions, 0, &count)
	return count
}

func solveNQueensHelper(positions []int, rowIndex int, count *int) {
	n := len(positions)
	if rowIndex == n {
		*count++
		return
	}
	for j := 0; j < n; j++ {
		ok := isValid(positions, rowIndex, j)
		if !ok {
			continue
		}
		positions[rowIndex] = j
		solveNQueensHelper(positions, rowIndex+1, count)
		positions[rowIndex] = -1
	}
}

func isValid(positions []int, rowIndex, colIndex int) bool {
	for i := 0; i < rowIndex; i++ {
		// 在同一列
		if positions[i] == colIndex {
			return false
		}
		// 在同一对角线
		if abs(i-rowIndex) == abs(positions[i]-colIndex) {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
