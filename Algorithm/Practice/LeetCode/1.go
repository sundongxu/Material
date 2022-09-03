package LeetCode

/**
 * 1. Two Sum
 * 描述：
 * 难度：Easy
 * 类型：Array
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[target-num] = i
	}
	for i, num := range nums {
		v, ok := m[num]
		if ok {
			return []int{v, i}
		}
	}
	return []int{}
}
