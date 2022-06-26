package LeetCode

import "math"

// 暴力枚举，超时，时间复杂度O(n^2)
//func maxProfit(prices []int) int {
//	max := math.MinInt
//	for i := 0; i < len(prices); i++ {
//		for j := i + 1; j < len(prices); j++ {
//			profit := prices[j] - prices[i]
//			if profit > max {
//				max = profit
//			}
//		}
//	}
//	if max < 0 {
//		return 0
//	}
//	return max
//}

// 贪心
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := math.MinInt
	curMin := prices[0]
	for i := 1; i < len(prices); i++ {
		profit = max(prices[i]-curMin, profit)
		curMin = min(prices[i], curMin)
	}
	return profit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 单调栈
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}
