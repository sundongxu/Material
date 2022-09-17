package LeetCode

/**
 * 152. Maximum Product Subarray
 * 描述：
 * 难度：Medium
 * 类型：DP
 */
// 动态规划
// 令fmax[i]表示以第i个元素结尾的乘积最大子数组的乘积
// 转移方程为：fmax[i] = max(fmax[i-1] * nums[i], fmin[i-1] * nums[i], nums[i])
// 令fmin[i]表示以第i个元素结尾的乘积最小子数组的乘积，考虑负数情况
// 转移方程为：fmin[i] = min(fmin[i-1] * nums[i], fmax[i-1] * nums[i], nums[i])
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
