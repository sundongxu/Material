package LeetCode

/**
 * 130. Surrounded Regions
 * 描述：
 * 难度：Medium
 * 类型：Matrix & DFS
 */
// 注意到题目解释中提到：任何边界上的 O 都不会被填充为 X。
// 可以想到，所有的不被包围的 O 都直接或间接与边界上的 O 相连。
// 可以利用这个性质判断 O 是否在边界上，具体地说：
// 对于每一个边界上的 O，以它为起点，标记所有与它直接或间接相连的字母 O；
// 最后遍历这个矩阵，对于每一个字母：
//（1）如果该字母被标记过，则该字母为没有被字母 X 包围的字母 O，我们将其还原为字母 O；
//（2）如果该字母没有被标记过，则该字母为被字母 X 包围的字母 O，我们将其修改为字母 X。

// 法一：深度优先搜索
var n, m int

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		// 以第一列元素为起点，标记所有与它直接或间接相连的字母O
		dfs(board, i, 0)
		// 以最后一列元素为起点
		dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		// 以第一行元素为起点，标记所有与它直接或间接相连的字母O
		dfs(board, 0, i)
		// 以最后一行元素为起点
		dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				// 被标记，一定是与边界O直接或间接相连的O，因为无法被全包围，所以将其还原为O
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				// 未被标记，且仍然为O，说明原来就是O，且未与边界O直接或间接相连，那么被全包围，改为X
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'  // 标记，说明原本是O，且是与边界O通过某条路径可以相邻的O，后续继续保持为O
	dfs(board, x+1, y) // 向右
	dfs(board, x-1, y) // 向左
	dfs(board, x, y+1) // 向上
	dfs(board, x, y-1) // 向下
}

// 法二：广度优先搜索，即层次遍历
var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	queue := make([][]int, 0)
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
			board[i][0] = 'A'
		}
		if board[i][m-1] == 'O' {
			queue = append(queue, []int{i, m - 1})
			board[i][m-1] = 'A'
		}
	}
	for i := 1; i < m-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
			board[0][i] = 'A'
		}
		if board[n-1][i] == 'O' {
			queue = append(queue, []int{n - 1, i})
			board[n-1][i] = 'A'
		}
	}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1] // 横纵坐标
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx < 0 || my < 0 || mx >= n || my >= m || board[mx][my] != 'O' {
				continue
			}
			queue = append(queue, []int{mx, my})
			board[mx][my] = 'A'
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
