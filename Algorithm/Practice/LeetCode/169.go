package LeetCode

/**
 * 169. Majority Element
 * 描述：
 * 难度：Easy
 * 类型：Array
 */
// 要求O(1)空间复杂度 -> 摩尔投票算法
// 摩尔投票算法基于如下事实：
// 	   每次从序列里选择两个不相同的数字删除掉（或称为“抵消”），最后剩下一个数字或几个相同的数字，就是出现次数大于总数一半的那个。
func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if res == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return res
}
