package LeetCode

/**
 * 134. Gas Station
 * 描述：
 * 难度：Medium
 * 类型：Greedy
 */
// 法一：用一个标志位j表示起点的前一个位置，每次尝试从位置j+1出发，如果发现不足以绕一圈，就重置j
// 如果gas之和 < cost之和，一定不能绕一圈，反之一定能
func canCompleteCircuit(gas []int, cost []int) int {
	total, sum, i, j := 0, 0, 0, -1
	for i < len(gas) {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if sum < 0 {
			// 说明以j+1为起点是不可行的，重置j到当前i，下一轮循环从j+1开始
			j = i
			sum = 0
		}
	}
	if total >= 0 {
		return j + 1
	}
	return -1
}

// 法二：
func canCompleteCircuit(gas []int, cost []int) int {
	// 从头到尾遍历每个加油站，并且检查以该加油站为起点，能否行驶一周
	for i, n := 0, len(gas); i < n; {
		// cnt记录以i为起点，能走过几个站点
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			// 以i为起点，能走过所有站点，即可绕一圈，返回起点i
			return i
		} else {
			// 否则从j+1=i+cnt+1为起点，重新尝试
			// 位置j=i+cnt的gas与cost之差一定是为负的，否则不会导致无法走到j+1=i+cnt+1位置
			// 推论：从[i, i+cnt]区间中的任何一个位置开始走，都无法走到位置i+cnt+1
			// 证明：
			// 位置i的gas与cost之差一定是为正的，即有油盈余的，否则从位置i无法走到位置i+1
			// 同理，从位置i出发，一直走到位置j=i+cnt，每到一个位置都是有油盈余的(>=0)
			// 所以前面走过的位置带给当前位置的都是正向效应，即永远不会带来一个负的油盈余，因为一旦出现负盈余，早就无法继续往后走了
			// 而位置j=i+cnt的gas与cost之差为负，且耗光了此前所有积累的油盈余，说明负得比较多
			// 那么不难得出，从[i, i+cnt]区间中的任何一个位置开始走，到位置j=i+cnt时，最终也会耗尽所有油盈余，所以直接把这区间给排除了
			// 所以下次尝试要从位置i+cnt+1开始走
			i += cnt + 1
		}
	}
	return -1
}
