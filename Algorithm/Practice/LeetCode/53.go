package LeetCode

// 数组中包含负数，找出和最大的子序列
// 暴力枚举法超时，时间复杂度O(n^2)，得想想有没有O(n)的解法，比如动态规划？
//func MaxSubArray(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	maxSum := math.MinInt
//	for i := 0; i < len(nums); i++ {
//		// 以i为子序列起点
//		if nums[i] > maxSum {
//			maxSum = nums[i]
//		}
//		if nums[i] < 0 {
//			// 起点即小于0，对求最大和毫无贡献，直接跳过
//			continue
//		}
//		sum := nums[i]
//		for j := i + 1; j < len(nums); j++ {
//			if sum+nums[j] < 0 {
//				// [i,j]之间子序列之和小于0，对求最大和毫无贡献，直接放弃
//				break
//			}
//			if sum > maxSum {
//				maxSum = sum
//			}
//			sum += nums[j]
//		}
//		if sum > maxSum {
//			maxSum = sum
//		}
//	}
//	return maxSum
//}

// 动态规划
// 对于任意一个元素，它有两种选择，第一种是加入之前的子序列中，第二种则是不加入
// 假设f[i]代表以第i个元素结尾的子序列的最大子序列和，那么最终答案就是数组f中找到最大元素，用一个变量随时记录即可，且f[0]=nums[0]
// 状态转移方程为：
// (1) 当f[i-1] > 0时，将nums[i]加入前面的子序列，则f[i] = f[i-1] + nums[i]
// (2) 当f[i-1] <= 0时，以nums[i]开始新的子序列，f[i] = nums[i]
// => f[i] = max(f[i-1]+nums[i], nums[i])
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// 模拟dp算法行为，本质上就只是一次对nums的遍历，记录一些中间状态即可
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
