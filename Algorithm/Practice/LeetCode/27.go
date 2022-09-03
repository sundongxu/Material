package LeetCode

/**
 * 27. Remove Element
 * 描述：
 * 难度：Easy
 * 类型：Array
 */
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
