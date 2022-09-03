package LeetCode

import "sort"

/**
 * 39. Combination Sum
 * 描述：
 * 难度：Medium
 * 类型：backtracking
 */
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	res, combination := make([][]int, 0), make([]int, 0)
	sort.Ints(candidates)
	findCombinationSum(candidates, combination, 0, target, &res)
	return res
}

func findCombinationSum(candidates []int, combination []int, sum int, target int, res *[][]int) {
	if sum == target {
		combinationCopy := make([]int, len(combination))
		copy(combinationCopy, combination)
		*res = append(*res, combinationCopy)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		if len(combination) > 0 {
			lastNum := combination[len(combination)-1]
			if lastNum > candidates[i] {
				continue
			}
		}
		combination = append(combination, candidates[i])
		findCombinationSum(candidates, combination, sum+candidates[i], target, res)
		combination = combination[:len(combination)-1]
	}
}
