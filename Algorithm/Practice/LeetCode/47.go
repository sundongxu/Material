package LeetCode

/**
 * 47. Permutations II
 * 描述：
 * 难度：Medium
 * 类型：Backtracking
 */
func PermuteUnique(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	uniqueNums := make([]int, 0)
	for k, _ := range counter {
		uniqueNums = append(uniqueNums, k)
	}

	res, permutation, pCounter := make([][]int, 0), make([]int, 0), make(map[int]int)
	findPermuteUnique(len(nums), uniqueNums, counter, permutation, pCounter, &res)
	return res
}

func findPermuteUnique(n int, uniqueNums []int, counter map[int]int, permutation []int, pCounter map[int]int, res *[][]int) {
	if len(permutation) == n {
		permutationCopy := make([]int, len(permutation))
		copy(permutationCopy, permutation)
		*res = append(*res, permutationCopy)
		return
	}
	for i := 0; i < len(uniqueNums); i++ {
		if counter[uniqueNums[i]] > pCounter[uniqueNums[i]] {
			permutation = append(permutation, uniqueNums[i])
			pCounter[uniqueNums[i]]++
			findPermuteUnique(n, uniqueNums, counter, permutation, pCounter, res)
			permutation = permutation[:len(permutation)-1]
			pCounter[uniqueNums[i]]--
		}
	}
}
