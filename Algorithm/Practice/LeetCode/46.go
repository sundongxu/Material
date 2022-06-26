package LeetCode

func Permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	permutation := make([]int, 0)
	findPermute(nums, permutation, &res)
	return res
}

func findPermute(nums []int, permutation []int, res *[][]int) {
	if len(permutation) == len(nums) {
		permutationCopy := make([]int, len(permutation))
		copy(permutationCopy, permutation)
		*res = append(*res, permutationCopy)
		return
	}
	counter := make(map[int]bool)
	for _, n := range permutation {
		counter[n] = true
	}
	for i := 0; i < len(nums); i++ {
		if !counter[nums[i]] {
			permutation = append(permutation, nums[i])
			findPermute(nums, permutation, res)
			permutation = permutation[:len(permutation)-1]
		}
	}
}
