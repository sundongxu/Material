package LeetCode

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
