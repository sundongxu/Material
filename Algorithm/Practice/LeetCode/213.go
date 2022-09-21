package LeetCode

/**
 * 213. House Robber II
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
// 动态规划，房屋首尾成环，要求不能抢劫相邻房屋
// 如果房屋数量大于两间，就必须考虑首尾相连的问题，第一间房屋和最后一间房屋不能同时偷窃
// 如何才能保证第一间房屋和最后一间房屋不同时偷窃呢？
//（1）如果偷窃了第一间房屋，则不能偷窃最后一间房屋，因此偷窃房屋的范围是第一间房屋到最后第二间房屋
//（2）如果偷窃了最后一间房屋，则不能偷窃第一间房屋，因此偷窃房屋的范围是第二间房屋到最后一间房屋
// 假设数组 nums 的长度为 n。
//（1）如果不偷窃最后一间房屋，则偷窃房屋的下标范围是 [0,n−2]
//（2）如果不偷窃第一间房屋，则偷窃房屋的下标范围是 [1, n-1]
// 在确定偷窃房屋的下标范围之后，即可用第 198 题的方法解决
// 对于两段下标范围分别计算可以偷窃到的最高总金额，其中的最大值即为在 n 间房屋中可以偷窃到的最高总金额
// 假设偷窃房屋的下标范围是 [start,end]，用 f[i] 表示在下标范围 [start,i] 内可以偷窃到的最高总金额，
// 那么就有如下的状态转移方程： f[i]=max(f[i−2] + nums[i], f[i−1])
// 边界条件为：
// 只有一间房屋，则偷窃该房屋：f[start] = nums[start]
// 只有两间房屋，偷窃其中金额较高的房屋：f[start+1]=max(nums[start],nums[start+1])
// 计算得到 f[end] 即为下标范围 [start,end] 内可以偷窃到的最高总金额。
// 分别取 (start,end)=(0,n−2) 和 (start,end)=(1,n−1) 进行计算，取两个 f[end] 中的最大值，即可得到最终结果。
func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
