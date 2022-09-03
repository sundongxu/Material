package LeetCode

/**
 * 200. Number of Islands
 * 描述：
 * 难度：Medium
 * 类型：Matrix
 */
// 解法一
var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res, visited := 0, make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				searchIslands(grid, &visited, i, j)
				res++
			}
		}
	}
	return res
}

func searchIslands(grid [][]byte, visited *[][]bool, x, y int) {
	(*visited)[x][y] = true
	for i := 0; i < 4; i++ {
		// 4个方向搜索
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(grid, nx, ny) && !(*visited)[nx][ny] && grid[nx][ny] == '1' {
			searchIslands(grid, visited, nx, ny)
		}
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

// 解法二：
func searchAt(grid [][]byte, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] == '0' {
		return 0
	}
	// 如果当前点(x,y)为陆地
	// 那么从上下左右四个方向出发，通过深度搜索即dfs的方式访问所有陆地(1)，遇到海洋(0)就结束遍历
	// 并将同一个岛屿的全部点都标记为0，此过程并不增加岛屿计数
	grid[x][y] = '0'       // 表示已经访问过，避免重复访问 => 核心逻辑
	searchAt(grid, x-1, y) // 向左一格
	searchAt(grid, x, y-1) // 向上一格
	searchAt(grid, x+1, y) // 向右一格
	searchAt(grid, x, y+1) // 向下一格
	return 1
}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			res += searchAt(grid, i, j)
		}
	}
	return res
}
