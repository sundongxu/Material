package LeetCode

/**
 * 75. Sort Colors
 * 描述：
 * 难度：Medium
 * 类型：Two Pointer
 */
func sortColors(nums []int) {
	zeroIndex, oneIndex := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n == 1 {
			nums[oneIndex] = 1
			oneIndex++
		} else if n == 0 {
			nums[zeroIndex] = 0
			zeroIndex++
		}
	}
}
