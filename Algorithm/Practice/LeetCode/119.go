package LeetCode

/**
 * 119. Pascal's Triangle II
 * 描述：
 * 难度：Easy
 * 类型：Array & Math
 */
func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	res := make([][]int, rowIndex+1)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
	}
	for _, row := range res {
		row[0] = 1
		row[len(row)-1] = 1
	}
	for i := 1; i < numRows; i++ {
		for j := 1; j < len(res[i])-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res[rowIndex]
}
