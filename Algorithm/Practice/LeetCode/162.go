package LeetCode

import (
	"math"
	"math/rand"
)

/**
 * 162. Find Peak Element
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
// 要求O(log n)时间复杂度，想到二分
// 且nums[-1] = nums[n] = -∞，即数组元素一定严格大于越界元素
func findPeakElement(nums []int) int {
	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, len(nums)-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

// 模拟爬坡
// 俗话说「人往高处走，水往低处流」。如果我们从一个位置开始，不断地向高处走，那么最终一定可以到达一个峰值位置
// 因此，我们首先在 [0, n)[0,n) 的范围内随机一个初始位置 ii，随后根据nums[i−1],nums[i],nums[i+1]三者的关系决定向哪个方向走：
// 如果nums[i−1]<nums[i]>nums[i+1]，那么位置i就是峰值位置，我们可以直接返回i作为答案；
// 如果nums[i−1]<nums[i]<nums[i+1]，那么位置i处于上坡，我们需要往右走，即i←i+1；
// 如果nums[i−1]>nums[i]>nums[i+1]，那么位置i处于下坡，我们需要往左走，即i←i−1；
// 如果nums[i−1]>nums[i]<nums[i+1]，那么位置i位于山谷，两侧都是上坡，我们可以朝任意方向走
// 如果我们规定对于最后一种情况往右走，那么当位置i不是峰值位置时：
// 如果nums[i]<nums[i+1]，那么我们往右走
// 如果nums[i]>nums[i+1]，那么我们往左走
func findPeakElement(nums []int) int {
	n := len(nums)
	idx := rand.Intn(n)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}

func findPeakElement1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > nums[mid+1] {
			// 如果 mid 较大，则左侧存在峰值，high = mid
			high = mid
		} else {
			// 如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
			low = mid + 1
		}
	}
	return low
}
