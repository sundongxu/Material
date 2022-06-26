package LeetCode

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) {
	positions, find := make([]position, 0), false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				// positions记录所有空白位置
				positions = append(positions, position{x: i, y: j})
			}
		}
	}
	putSudoku(&board, positions, 0, &find)
}

func putSudoku(board *[][]byte, positions []position, posIndex int, find *bool) {
	if *find {
		return
	}
	if posIndex == len(positions) {
		*find = true
		return
	}
	// i为要填入的数字，范围[1,9]
	for i := 1; i <= 9; i++ {
		// 检查是否可以放入i
		if checkSudoku(board, positions[posIndex], i) && !*find {
			// 可以放入i, 则放入i
			(*board)[positions[posIndex].x][positions[posIndex].y] = byte(i) + '0'
			putSudoku(board, positions, posIndex+1, find)
			if *find {
				return
			}
			(*board)[positions[posIndex].x][positions[posIndex].y] = '.'
		}
	}
}

// 检查把num填入pos位置是否合法
func checkSudoku(board *[][]byte, pos position, num int) bool {
	// 检查pos所在横行是否能放num
	for i := 0; i < 9; i++ {
		if (*board)[pos.x][i] != '.' && int((*board)[pos.x][i]-'0') == num {
			return false
		}
	}

	// 检查纵列
	for j := 0; j < 9; j++ {
		if (*board)[j][pos.y] != '.' && int((*board)[j][pos.y]-'0') == num {
			return false
		}
	}

	// 检查九宫格 => 确定(x,y)这个点所在九宫格范围如何表示
	// x范围：[x/3*3, x/3*3+2] => [x-x%3, x-x%3+3)
	// y起点：[y/3*3, y/3*3+2] => [y-y%3, y-y%3+3)
	startX, startY := pos.x/3*3, pos.y/3*3
	for i := startX; i < startX+3; i++ {
		for j := startY; j < startY+3; j++ {
			if (*board)[i][j] != '.' && int((*board)[i][j]-'0') == num {
				return false
			}
		}
	}

	return true
}

// 单引号，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。
// 双引号，才是字符串，实际上是字符数组。可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。
// 反引号，表示字符串字面量，但不支持任何转义序列。字面量 raw literal string 的意思是，你定义时写的啥样，它就啥样，你有换行，它就换行。你写转义字符，它也就展示转义字符。
