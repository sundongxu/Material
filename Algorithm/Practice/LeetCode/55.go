package LeetCode

/**
 * 55. Jump Game
 * 描述：
 * 难度：Medium
 * 类型：DP & Greedy
 */
// 贪心
// 推论：如果能跳到最后一级，那么就能跳到任意一级
// 分析：如果能跳到最后一级，那么只有两种情况
// 第一种是从倒数第二级开始跳且只跳一级，这种情况假设了一定能跳到倒数第二级，且从倒数第二级还能够跳至少1级，在推论前提下，能够到最倒数第二级
// 第二种是从倒数第二级之前的位置开始跳，且跳不止一级，在推论前提下，既然都能够跳到最后一级了，那么一定也能少跳一级，从而跳到倒数第二级
// 在上述两种可能中，在一定能跳到最后一级的推论前提下，都能推导出一定能跳到倒数第二级，以此类推，一定能跳到倒数第三级、倒数第四级...第一级，从而可以跳到任意一级，于是推论正确
func canJump(nums []int) bool {
	maxJump := 0
	for index, num := range nums {
		if index > maxJump {
			// 前面的步骤所能跳至的最大级数，都到不了index位置了，说明在此位置发生了中断了，故而不能到终点，直接返回
			return false
		}
		// 前面的步骤能到达位置index，则可以从位置index再计算能跳到的最远距离
		maxJump = max(maxJump, index+num)
	}
	return maxJump >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 反向贪心，从高往低下楼梯，如果能正向上至任意一级，那么也应该能反向下至任意一级
func canJump2(nums []int) bool {
	curStep := len(nums) - 1 // 从终点即最后一级开始往左下楼梯
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= curStep {
			// 从位置i出发，能往左最远下到位置curStep
			curStep = i
		}
	}
	return curStep == 0
}

// 动态规划
// 假设f[i]表示，到达第i个位置时，剩余还能跳的最大级数
// 能够到达第i个位置，包含两种情况
// 第一种情况：从第i-1个位置开始跳，能跳的级数是nums[i-1]，则达到第i个位置时，还剩余的最大级数是nums[i-1]-1
// 第二种情况：从第i-1个位置以前的某个位置开始跳，如果能达到第i个位置，说明也能通过少跳一级从而到达第i-1个位置，到达第i-1个位置时剩余还能跳的最大级数是f[i-1]，那么再多跳一级到达第i个位置时，还剩余的最大级数是f[i-1]-1
// 状态转移方程为：f[i] = max(f[i-1], nums[i-1]) - 1
// 那么只要有f[i-1]这一个前置状态，就足够算出当前状态f[i]
// 初始状态即f[0]，代表含义是到达第0个位置时，剩余还能跳的最大级数，一开始就是从第0级开始，不需要跳，所以f[0]=0
// 如果中途有任何一个f[k]<0，说明到达第k个位置时，剩余一级都不能跳了，那么也无法到达>k级的台阶，从而无法到达最后一级，直接返回false即可
// 最终就只要看最后一级台阶对应的f[len(nums)-1]的值，如果f[len(nums)-1]>=0，说明可以到达最后一级
func canJump3(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i-1]) - 1
		if dp[i] < 0 {
			// 说明最远都到不了第i个位置，从而也就到不了最后一级
			return false
		}
	}
	return dp[len(nums)-1] >= 0
}
