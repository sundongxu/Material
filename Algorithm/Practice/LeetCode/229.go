package LeetCode

/**
 * 229. Majority Element II
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
// 要求O(1)空间复杂度
func majorityElement(nums []int) []int {
	// since we are checking if a num appears more than 1/3 of the time
	// it is only possible to have at most 2 nums (>1/3 + >1/3 = >2/3)
	count1, count2, candidate1, candidate2 := 0, 0, 0, 0
	// Select Candidates
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			// We have a bad first candidate, replace!
			candidate1, count1 = num, 1
		} else if count2 == 0 {
			// We have a bad second candidate, replace!
			candidate2, count2 = num, 1
		} else {
			// Both candidates suck, boo!
			count1--
			count2--
		}
	}
	// Recount!
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}
	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}
	if count1 > length/3 {
		return []int{candidate1}
	}
	if count2 > length/3 {
		return []int{candidate2}
	}
	return []int{}
}
