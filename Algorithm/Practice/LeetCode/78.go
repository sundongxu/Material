package LeetCode

/**
 * 78. Subsets
 * 描述：
 * 难度：Medium
 * 类型：Backtracking
 */
//func Subsets(nums []int) [][]int {
//	if len(nums) == 0 {
//		return [][]int{}
//	}
//	res, subset := make([][]int, 0), make([]int, 0)
//	for i := 0; i <= len(nums); i++ {
//		// generate subset containing i numbers
//		findSubsets(nums, subset, 0, i, &res)
//	}
//	return res
//}
//
//func findSubsets(nums []int, subset []int, index int, n int, res *[][]int) {
//	if len(subset) == n {
//		subsetCopy := make([]int, len(subset))
//		copy(subsetCopy, subset)
//		*res = append(*res, subsetCopy)
//		return
//	}
//	for i := index; i < len(nums); i++ {
//		subset = append(subset, nums[i]) // choose nums[i]
//		findSubsets(nums, subset, i+1, n, res)
//		subset = subset[:len(subset)-1] // or not => backtrack
//	}
//}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, subset := make([][]int, 0), make([]int, 0)
	findSubsets(nums, subset, 0, &res)
	return res
}

func findSubsets(nums []int, subset []int, index int, res *[][]int) {
	if index == len(nums) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*res = append(*res, subsetCopy)
		return
	}
	// not choose nums[index]
	findSubsets(nums, subset, index+1, res)
	// choose nums[index]
	subset = append(subset, nums[index])
	findSubsets(nums, subset, index+1, res)
	subset = subset[:len(subset)-1]
}
