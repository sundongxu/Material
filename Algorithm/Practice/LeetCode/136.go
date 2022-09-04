package LeetCode

/**
 * 136. Single Number
 * 描述：
 * 难度：Easy
 * 类型：Array
 */
// 要求O(n)时间复杂度，O(1)空间复杂度
// 任何一个数字异或它自己都等于0，那么将数组全部元素异或一遍，剩下的就是
// 利用异或运算的 归零律：X⊕X = 0 和 恒等律：X⊕0 = X
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
