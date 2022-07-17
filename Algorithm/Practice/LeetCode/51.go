package LeetCode

import "strings"

// 51. N-Queens
// 同一行、同一列、同一斜线上的皇后都会自动攻击
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	positions := make([]int, n) // 第i个元素的值positions[i]表示，第i行第positions[i]列放置皇后
	for i, _ := range positions {
		positions[i] = -1
	}
	solveNQueensHelper(positions, &res, 0)
	return res
}

func solveNQueensHelper(positions []int, res *[][]string, rowIndex int) {
	n := len(positions)
	if rowIndex == n {
		solution := make([]string, 0)
		// 终止条件，也是收敛条件，意味着找到了一个可行解
		for i := 0; i < n; i++ {
			row := generateDefaultRow(n)
			for j := 0; j < n; j++ {
				if j == positions[i] {
					row[j] = "Q"
				}
			}
			solution = append(solution, strings.Join(row, ""))
		}
		*res = append(*res, solution)
		return
	}
	for j := 0; j < n; j++ {
		ok := isValid(positions, rowIndex, j)
		if !ok {
			continue
		}
		positions[rowIndex] = j
		solveNQueensHelper(positions, res, rowIndex+1)
		positions[rowIndex] = -1
	}
}

func generateDefaultRow(n int) []string {
	row := make([]string, n)
	for j, _ := range row {
		row[j] = "."
	}
	return row
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
