package LeetCode

import "math"

// 188. Best Time to Buy and Sell Stock IV
// 解法一：二维动规
// 与其余的股票问题类似，我们使用一系列变量存储「买入」的状态，再用一系列变量存储「卖出」的状态，通过动态规划的方法即可解决本题。
// 用 buy[i][j]表示对于数组prices[0..i]中的价格而言，进行恰好j笔交易，并且当前手上持有一支股票，这种情况下的最大利润
// 同理，用 sell[i][j]表示恰好进行j笔交易，并且当前手上不持有股票，这种情况下的最大利润。
// 对状态转移方程进行推导。对于buy[i][j]，我们考虑当前手上持有的股票是否是在第i天买入的。
// 如果是第i天买入的，那么在第i-1天时，我们手上不持有股票，对应状态sell[i-1][j]，并且需要扣除prices[i]的买入花费
// 如果不是第i天买入的，那么在第i-1天时，我们手上持有股票，对应状态buy[i−1][j]
// 那么可以得到状态转移方程：
// （1）buy[i][j]=max{buy[i−1][j],sell[i−1][j]−price[i]}
// 同理对于sell[i][j]，
// 如果是第i天卖出的，那么在第i−1天时，我们手上持有股票，对应状态buy[i−1][j−1]，并且需要增加prices[i] 的卖出收益
// 如果不是第i天卖出的，那么在第i-1天时，我们手上不持有股票，对应状态sell[i−1][j]
// 那么可以得到状态转移方程：
// （2）sell[i][j]=max{sell[i−1][j],buy[i−1][j−1]+price[i]}
// 由于在所有的n天结束后，手上不持有股票对应的最大利润一定是严格大于手上持有股票对应的最大利润的
// 然而完成的交易数并不是越多越好（例如数组prices单调递减，那么我们不进行任何交易才是最优的）
// 因此最终的答案即为sell[n−1][0..k]中的最大值
// 在上述的状态转移方程中，确定边界条件是非常重要的步骤。
// 考虑将所有的buy[0][0..k]以及sell[0][0..k]设置为边界
// 对于buy[0][0..k]，由于只有prices[0]唯一的股价，因此我们不可能进行过任何交易
// 那么我们可以将所有的buy[0][1..k]设置为一个非常小的值，表示不合法的状态
// 而对于buy[0][0]，它的值为prices[0]，即「我们在第0天以prices[0]的价格买入股票」是唯一满足手上持有股票的方法
// 对于sell[0][0..k]，同理我们可以将所有的sell[0][1..k]设置为一个非常小的值，表示不合法的状态
// 而对于sell[0][0]，它的值为0，即「我们在第0天不做任何事」是唯一满足手上不持有股票的方法
// 在设置完边界之后，就可以使用二重循环，在i∈[1,n),j∈[0,k]的范围内进行状态转移
// 需要注意的是，sell[i][j] 的状态转移方程中包含buy[i−1][j−1]，在j=0时其表示不合法的状态
// 因此在j=0时，我们无需对sell[i][j]进行转移，让其保持值为0即可。
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return max(sell[n-1]...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

// 解法二：一维动规：滚动数组
// 注意到在状态转移方程中，buy[i][j]和sell[i][j]都从buy[i−1][..]以及sell[i−1][..]转移而来，
// 因此我们可以使用一维数组而不是二维数组进行状态转移，即：
// b[j]←max{b[j],s[j]−price[i]}
// s[j]←max{s[j],b[j−1]+price[i]}
// 这样的状态转移方程会因为状态的覆盖而变得不同
// 例如如果我们先计算b而后计算s，那么在计算到s[j]时，其状态转移方程中包含的b[j−1]这一项的值已经被覆盖了
// 即本来应当是从二维表示中的buy[i−1][j−1]转移而来，而现在却变成了buy[i][j−1]
// 但其仍然是正确的。我们考虑buy[i][j−1]的状态转移方程：
// b[j−1]←buy[i][j−1]=max{buy[i−1][j−1],sell[i−1][j−1]−price[i]}
// 那么 s[j]s[j] 的状态转移方程实际上为：
// s[j]←max{s[j],buy[i−1][j−1]+prices[i],sell[i−1][j−1]}
// 为什么 s[j]s[j] 的状态转移方程中会出现sell[i−1][j−1]这一项？
// 实际上，我们是把「在第i天以prices[i]的价格买入，并在同一天以prices[i] 的价格卖出」看成了一笔交易
// 这样对应的收益即为：
// sell[i−1][j−1]−prices[i]+prices[i]
// 也就是sell[i−1][j−1] 本身。这种在同一天之内进行一笔交易的情况，收益为零
// 它并不会带来额外的收益，因此对最终的答案并不会产生影响，状态转移方程在本质上仍然是正确的
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64 / 2
		sell[i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j]-prices[i])
			sell[j] = max(sell[j], buy[j-1]+prices[i])
		}
	}
	return max(sell...)
}
