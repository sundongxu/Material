package LeetCode

import "math"

// 174. Dungeon Game
// 要求计算出骑士的最小初始生命值，即刚到左上角房间时候具有的生命值
// 生命值到0或小于0骑士就死了

// 考虑从右下往左上进行动态规划
// 令dp[i][j] 表示从坐标(i,j)到终点所需的最小初始值
// 换句话说，当我们到达坐标(i,j)时，如果此时我们的路径和不小于dp[i][j]，我们就能到达终点
// 这样一来，我们就无需担心路径和的问题，只需要关注最小初始值
// 对于dp[i][j]，我们只要关心dp[i][j+1]和dp[i+1][j]的最小值minn
// 记当前格子的值为dungeon(i,j)，那么在坐标(i,j)的初始值只要达到minn−dungeon(i,j) 即可
// 同时，初始值还必须大于等于1
// 这样就可以得到状态转移方程：
// dp[i][j]=max(min(dp[i+1][j],dp[i][j+1])−dungeon(i,j),1)
// 最终答案即为dp[0][0]。
// 边界条件为，当i=n−1或者j=m−1时，dp[i][j] 转移需要用到的dp[i][j+1]和dp[i+1][j]中有无效值，因此代码实现中给无效值赋值为极大值
//特别地，dp[n−1][m−1]转移需要用到的[m]dp[n−1][m]和dp[n][m−1]均为无效值，因此给这两个值赋值为1
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			minn := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minn-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
