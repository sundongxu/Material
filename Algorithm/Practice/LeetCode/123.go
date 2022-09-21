package LeetCode

/**
 * 123. Best Time to Buy and Sell Stock III
 * 描述：
 * 难度：Hard
 * 类型：DP
 */

// 最多完成两笔交易
// 动态规划
// 由于我们最多可以完成两笔交易，因此在任意一天结束之后，我们会处于以下五个状态中的一种：
//（1）未进行过任何操作；
//（2）只进行过一次买操作；
//（3）进行了一次买操作和一次卖操作，即完成了一笔交易；
//（4）在完成了一笔交易的前提下，进行了第二次买操作；
//（5）完成了全部两笔交易。
// 由于第一个状态的利润显然为0，因此可以不用将其记录。
// 对于剩下的四个状态，我们分别将它们的最大利润记为buy1、sell1、buy2、sell2
// 最好先用一维DP来理解，令buy1[i]表示第i天只进行了一次买操作的最大利润，其它同理
// 状态转移方程依次为：
// （1）buy1[i] = max(buy1[i-1], -price[i])
// （2）sell1[i] = max(sell1[i-1], buy1[i-1] + price[i])
// （3）buy2[i] = max(buy2[i-1], sell1[i-1] - price[i])
// （4）sell2[i] = max(sell2[i-1], buy2[i-1] + price[i])
// 实际可能的答案只会是0、sell1[i]、sell2[i]中的最大值，不可能是buy1或buy2，因为如果买了不卖不可能得到最大利润

// 再来考虑优化一维DP，用四个变量buy1、sell1、buy2、sell2分别代表上述四个数组
// 那么此时状态转移方程变为：
// （1）buy1 = max(buy1, -price[i])
// （2）sell1 = max(sell1, buy1 + price[i])
// （3）buy2 = max(buy2, sell1 - price[i])
// （4）sell2 = max(sell2, buy2 + price[i])
// 以（2）为例进行说明，此时等式右边的sell1其实是第i-1天的sell1值即上面的sell1[i-1]，这个不难理解
// 而用已经更新了第i天的buy1代替上面的buy1[i-1]，是因为此时buy1相比buy1[i-1]本质上就是多考虑了第i天的买入
// 但是考虑到buy1计算是通过max取了buy1[i-1]和-price[i]的最大值，所以显然buy1实际就等于buy1[i-1]，
// 所以这多考虑的一天实际不会对结果有任何影响，二者可以认为是等价的
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
