package LeetCode

/**
 * 81. Search in Rotated Sorted Array II
 * 描述：
 * 难度：Medium
 * 类型：Array & Binary Search & Recursion
 */
// rotated sorted array => binary search
// elements may not be unique
// 解法同33
func search(nums []int, target int) bool {
	// nums = [4,5,6,7,0,1,2]
	i := len(nums) - 2
	// 从右边开始，找一个结束升序排列的临界点nums[i]
	for i >= 0 && nums[i] <= nums[i+1] {
		i--
	}
	// nums[i]右边的序列一定是升序
	// nums[i]左边的序列也一定是升序
	// 问题变成，如何在两个升序数组中，找到目标值
	leftIndex := binarySearchInAscendingArray(nums, 0, i, target)
	rightIndex := binarySearchInAscendingArray(nums, i+1, len(nums)-1, target)
	return leftIndex || rightIndex
}

func binarySearchInAscendingArray(nums []int, start, end, target int) bool {
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return true
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return false
}

func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return false
}
