package LeetCode

import "sort"

func fourSum(nums []int, target int) [][]int {
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
		if (uniqNums[i]*4 == target) && counter[uniqNums[i]] >= 4 {
			results = append(results, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*3+uniqNums[j] == target) && counter[uniqNums[i]] > 2 {
				results = append(results, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*3+uniqNums[i] == target) && counter[uniqNums[j]] > 2 {
				results = append(results, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i]*2 == target) && counter[uniqNums[j]] > 1 && counter[uniqNums[i]] > 1 {
				results = append(results, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) && counter[uniqNums[i]] > 1 {
					results = append(results, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[j]*2+uniqNums[i]+uniqNums[k] == target) && counter[uniqNums[j]] > 1 {
					results = append(results, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[k]*2+uniqNums[i]+uniqNums[j] == target) && counter[uniqNums[k]] > 1 {
					results = append(results, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				left := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if left > uniqNums[k] && counter[left] > 0 {
					results = append(results, []int{uniqNums[i], uniqNums[j], uniqNums[k], left})
				}
			}
		}
	}
	return results
}
