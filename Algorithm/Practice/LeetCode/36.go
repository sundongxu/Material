package LeetCode

func isValidSudoku(board [][]byte) bool {
	// 检查横行是否合法
	for i := 0; i < 9; i++ {
		counterRow := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if num > 9 || num < 1 {
					return false
				}
				if counterRow[num] {
					return false
				}
				counterRow[num] = true
			}
		}
	}

	// 检查纵列是否合法
	for i := 0; i < 9; i++ {
		counterCol := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				num := int(board[j][i] - '0')
				if num > 9 || num < 1 {
					return false
				}
				if counterCol[num] {
					return false
				}
				counterCol[num] = true
			}
		}
	}

	// 检查九宫格是否合法
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			counterSquare := make(map[int]bool)
			// 横向第i个，纵向第j个，九宫格
			for m := i * 3; m < i*3+3; m++ {
				for n := j * 3; n < j*3+3; n++ {
					if board[m][n] != '.' {
						num := int(board[m][n] - '0')
						if num > 9 || num < 1 {
							return false
						}
						if counterSquare[num] {
							return false
						}
						counterSquare[num] = true
					}
				}
			}
		}
	}

	return true
}
