package LeetCode

/**
 * 198. House Robber
 * 描述：
 * 难度：Medium
 * 类型：DP
 */
// 动态规划，要求不能抢劫相邻的两个房屋
// 对于第 i (i>2) 间房屋，有两个选项：
// 偷窃第 i 间房屋，那么就不能偷窃第 i−1 间房屋，偷窃总金额为前 i−2 间房屋的最高总金额与第 i 间房屋的金额之和。
// 不偷窃第 i 间房屋，偷窃总金额为前 i−1 间房屋的最高总金额
// 令f[i]为前i间房屋能抢到的最高总金额
// 转移方程为：f[i] = max{f[i-2]+nums[i], f[i-1]}
// 边界条件为：f[0] = nums[0]，f[1] = max{nums[0], nums[1]}，最终答案为f[n-1]，其中n为nums数组的长度
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
