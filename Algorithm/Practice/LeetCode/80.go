package LeetCode

func removeDuplicates(nums []int) int {
	k, cnt := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
			cnt = 1
		} else if nums[k] == nums[i] && cnt < 2 {
			k++
			nums[k] = nums[i]
			cnt++
		}
	}
	return k + 1
}
