package LeetCode

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	counter := make(map[int]int)
	for _, num := range candidates {
		counter[num]++
	}

	uniqueCandidates := make([]int, 0)
	for k, _ := range counter {
		uniqueCandidates = append(uniqueCandidates, k)
	}

	sort.Ints(uniqueCandidates)
	res, combination, cCounter := make([][]int, 0), make([]int, 0), make(map[int]int)
	findCombinationSum2(uniqueCandidates, combination, 0, cCounter, counter, target, &res)
	return res
}

func findCombinationSum2(candidates []int, combination []int, sum int, cCounter map[int]int, counter map[int]int, target int, res *[][]int) {
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

		if cCounter[candidates[i]] == counter[candidates[i]] {
			continue
		}
		if len(combination) > 0 {
			lastNum := combination[len(combination)-1]
			if lastNum > candidates[i] {
				continue
			}
		}
		combination = append(combination, candidates[i])
		cCounter[candidates[i]]++
		findCombinationSum2(candidates, combination, sum+candidates[i], cCounter, counter, target, res)
		combination = combination[:len(combination)-1]
		cCounter[candidates[i]]--
	}
}
