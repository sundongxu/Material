package LeetCode

/**
 * 34. Find First and Last Position of Element in Sorted Array
 * 描述：
 * 难度：Medium
 * 类型：Array & Binary Search
 */
func SearchRange(nums []int, target int) []int {
	// nums = [5,7,7,8,8,10]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// nums[mid] == target
			low, high := mid, mid
			for low-1 >= 0 && nums[low-1] == target {
				low--
			}
			for high+1 < len(nums) && nums[high+1] == target {
				high++
			}
			return []int{low, high}
		}
	}
	return []int{-1, -1}
}
