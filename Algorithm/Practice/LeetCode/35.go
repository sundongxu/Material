package LeetCode

/**
 * 35. Search Insert Position
 * 描述：
 * 难度：Easy
 * 类型：Array & Binary Search
 */
func SearchInsert(nums []int, target int) int {
	// nums = [1,3,5,6]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// 可以排除nums[mid]
			right = mid - 1
		} else {
			// nums[mid] < target
			// 可以排除nums[mid]，因为mid位置原本元素比target小，要想插入target，一定在位置mid之后
			left = mid + 1
			if left == len(nums) || nums[left] >= target {
				return left
			}
		}
	}
	return 0
}
