package LeetCode

import "sort"

/**
 * 15. 3Sum
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
func ThreeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}
	counter := make(map[int]int)
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := make([]int, 0)
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	results := make([][]int, 0)

	sort.Ints(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		if uniqNums[i] == 0 && counter[uniqNums[i]] > 2 {
			results = append(results, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] > 1 {
				results = append(results, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if uniqNums[j]*2+uniqNums[i] == 0 && counter[uniqNums[j]] > 1 {
				results = append(results, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			left := 0 - uniqNums[i] - uniqNums[j]
			if left > uniqNums[j] && counter[left] > 0 {
				results = append(results, []int{uniqNums[i], uniqNums[j], left})
			}
		}
	}

	return results
}
