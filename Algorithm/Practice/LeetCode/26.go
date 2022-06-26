package LeetCode

// brutal force - time limit exceeded
//func RemoveDuplicates(nums []int) int {
//	k := len(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if i == k {
//			break
//		}
//		for nums[i] == nums[i+1] {
//			// push nums[i] to the end of array
//			cur := nums[i]
//			for j := i; j < len(nums)-1; j++ {
//				nums[j] = nums[j+1]
//			}
//			nums[len(nums)-1] = cur
//			k--
//		}
//	}
//	return k
//}

// do not care what left except the first k elements
func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
