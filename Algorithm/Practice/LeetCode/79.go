package LeetCode

/**
 * 79. Word Search
 * 描述：
 * 难度：Medium
 * 类型：Matrix & DFS
 */
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	// 递归函数参数
	res := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = findWord(board, m, n, i, j, word)
			if res {
				break
			}
		}
	}
	return res
}

func findWord(board [][]byte, m, n, i, j int, word string) bool {
	if i >= m || j >= n || i < 0 || j < 0 {
		return false
	}
	if len(word) == 0 || (len(word) == 1 && board[i][j] == word[0]) {
		// 递归结束条件
		return true
	}
	if board[i][j] != word[0] {
		return false
	}
	return findWord(board, m, n, i-1, j, word[1:]) || findWord(board, m, n, i+1, j, word[1:]) || findWord(board, m, n, i, j-1, word[1:]) || findWord(board, m, n, i, j+1, word[1:])
}
