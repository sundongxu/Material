package LeetCode

/**
 * 153. Find Minimum in Rotated Sorted Array
 * 描述：
 * 难度：Medium
 * 类型：Array & Binary Search
 */
// 要求时间复杂度O(log n)，旋转数组升序，且元素唯一
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			// 说明mid一定在最小值右侧，或本身就是最小值，所以right可以等于mid
			right = mid
		} else {
			// 说明mid一定在最小值左侧，由于它大于（无重复元素所以无法等于）right，所以一定不可能是最小元素故mid是一定要排除的，所以left不能等于mid
			left = mid + 1
		}
	}
	return nums[left]
}
