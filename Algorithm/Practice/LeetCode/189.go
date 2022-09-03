package LeetCode

/**
 * 189. Rotate Array
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
// 空间复杂度O(n)
// nums数组的第i个元素，右移k位后的位置为：(i+k) % len(nums)
//func rotate(nums []int, k int) {
//	newNums := make([]int, len(nums))
//	for i, v := range nums {
//		newNums[(i+k)%len(nums)] = v
//	}
//	copy(nums, newNums)
//}

// 空间复杂度O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(arr []int) {
	for i, n := 0, len(arr); i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}
