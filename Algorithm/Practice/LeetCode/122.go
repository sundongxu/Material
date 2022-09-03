package LeetCode

/**
 * 122. Best Time to Buy and Sell Stock II
 * 描述：
 * 难度：Medium
 * 类型：Greedy
 */
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	cur, res := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > cur {
			// 存在利润，累加利润
			res += prices[i] - cur
		}
		// 可能是当天卖出，当天买进 => if条件满足时
		// 可能是当天是更低价，所以当天才买进  => if条件不满足时
		cur = prices[i]
	}
	return res
}

// 贪心法，总是低进高出，把相邻的所有价格差价相加即所求
func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}
