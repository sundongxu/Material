package LeetCode

import "sort"

// what about find previous permutations ?
func nextPermutation(nums []int) {
	// nums=[1,2,4,3]
	// 从右边开始，找到一个降序排列的终结点nums[i]，它比nums[i+1]要大
	// 为什么要找降序排列？=> 因为降序排列说明这个子排列是以nums[i]作为首元素，用num[i]及其之后所有元素所能构成的最大排列
	// 下一个排列最好找，一定不是以nums[i]开头，而是以其右边序列中比它大的元素中的最小元素为开头的，且首元素之后的序列应该是以该元素为首的最小序列，即升序序列
	// 还是从右边开始，找到这个在nums[i]右边序列中比它大的元素中的最小元素，并将其和nums[i]互换位置，然后再将位置i之后的序列排序得到升序序列即可
	// 而由于互换位置后，位置i之后的序列仍为降序序列，要得到升序序列，除了直接原地排序，也可以直接反转此序列
	// 故nums[i]的右边一定为降序排列
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 如果nums本身就是降序排列，那么i==-1，此时访问nums[i]就会越界

	if i >= 0 {
		// 在nums[i]的右边序列中找到比它大的数中最小的数作为nums[j]
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// nums[i]=2, nums[j]=3
		// [1,2,4,3] => 互换nums[i]和nums[j] => [1,3,4,2]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 直接按升序对第i个元素之后的序列进行重新排序
	// [1,3,4,2] => 对nums[i]右边的序列重新按升序排序 => [1,3,2,4]
	sort.Ints(nums[i+1:])

	// 由于互换位置后，第i个元素之后的序列仍为降序排列，要得到升序排列，直接反转该序列即可
	// [1,3,4,2] => 反转nums[i]右边的序列 => [1,3,2,4]
	//p := i + 1
	//q := len(nums) - 1
	//for p < q {
	//	nums[p], nums[q] = nums[q], nums[p]
	//	p++
	//	q--
	//}
}
