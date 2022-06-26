package LeetCode

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
