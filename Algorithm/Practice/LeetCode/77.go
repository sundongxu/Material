package LeetCode

/**
 * 77. Combinations
 * 描述：
 * 难度：Medium
 * 类型：Backtracking
 */
//func Combine(n int, k int) [][]int {
//	nums := make([]int, 0)
//	i := 1
//	for i <= n {
//		nums = append(nums, i)
//		i++
//	}
//	if n == k {
//		return [][]int{nums}
//	}
//	res := make([][]int, 0)
//	//findCombine(&nums, k, []int{}, &res)
//	return res
//}

//func findCombine(nums *[]int, leftNum int, combination []int, res *[][]int) {
//	if leftNum == 0 {
//		*res = append(*res, combination)
//		return
//	}
//	diff := getDiff(nums, &combination)
//	for _, num := range *diff {
//		lastNum := 0
//		if len(combination) > 0 {
//			lastNum = combination[len(combination)-1]
//		}
//		if num > lastNum {
//			combinationCopy := make([]int, len(combination))
//			copy(combinationCopy, combination)
//			combinationCopy = append(combinationCopy, num)
//			findCombine(nums, leftNum-1, combinationCopy, res)
//		}
//	}
//}
//
//func getDiff(num1 *[]int, num2 *[]int) *[]int {
//	diff := make([]int, 0)
//	counter := make(map[int]bool)
//	for _, num := range *num2 {
//		counter[num] = true
//	}
//	for _, num := range *num1 {
//		if !counter[num] {
//			diff = append(diff, num)
//		}
//	}
//	return &diff
//}

func Combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	res := make([][]int, 0)
	nums := make([]int, 0)
	findCombine(n, k, 1, nums, &res)
	return res
}

func findCombine(n, k, start int, nums []int, res *[][]int) {
	if len(nums) == k {
		combination := make([]int, len(nums))
		copy(combination, nums)
		*res = append(*res, combination)
		return
	}
	leftNum := k - len(nums)
	// assume n = 5, k = 3
	// current combination size: 【k-leftNum】
	// current select 1, future select 【leftNum - 1】, must prepare a set of size 【leftNum - 1】
	// current select 1 start from 【start】, in order to reserve a set of size 【leftNum - 1】, the maximum start will be 【n - (leftNum - 1)】= 【n-leftNum+1】
	// select range from 【start, n - leftNum + 1】
	for i := start; i <= n-leftNum+1; i++ {
		nums = append(nums, i)
		findCombine(n, k, i+1, nums, res)
		nums = nums[:len(nums)-1]
	}
}
