package LeetCode

// rotated sorted array => binary search
// all elements are unique
func search1(nums []int, target int) int {
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
	if leftIndex != -1 {
		return leftIndex
	}
	if rightIndex != -1 {
		return rightIndex
	}
	return -1
}

func binarySearchInAscendingArray(nums []int, start, end, target int) int {
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			// 如果满足这一条件，那么根据旋转升序数组的性质，[left, mid]一定为单调递增区间，判断target是否在这区间
			if nums[left] <= target && target <= nums[mid] {
				// 若target位于[left, mid]区间，那么mid右边区域可以排除
				right = mid
			} else {
				// 否则，mid左边区域可以排除
				left = mid + 1
			}
		} else {
			// 如果[left,mid]不是单调递增区间，那么根据旋转升序数组的性质，[mid, right]一定为单调递增区间，判断target是否在这区间
			if nums[mid] <= target && target <= nums[right] {
				// 若target位于[mid, right]区间，那么mid左边区域可以排除
				left = mid
			} else {
				// 否则，mid右边区域可以排除
				right = mid - 1
			}
		}
	}
	return -1
}

func search3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
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
	return -1
}
