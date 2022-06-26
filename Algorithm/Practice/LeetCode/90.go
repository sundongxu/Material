package LeetCode

func SubsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	uniqueNums := make([]int, 0)
	for k, _ := range counter {
		uniqueNums = append(uniqueNums, k)
	}

	res, subset := make([][]int, 0), make([]int, 0)
	findSubsetsWithDup(uniqueNums, counter, 0, subset, &res)
	return res
}

func findSubsetsWithDup(nums []int, counter map[int]int, index int, subset []int, res *[][]int) {
	if index == len(nums) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*res = append(*res, subsetCopy)
		return
	}
	// choose nums[index]
	if counter[nums[index]] > 0 {
		subset = append(subset, nums[index])
		counter[nums[index]]--
		findSubsetsWithDup(nums, counter, index, subset, res) // key, next Recursion( still can process nums[index]
		counter[nums[index]]++
		subset = subset[:len(subset)-1]
	}
	// not choose nums[index]
	findSubsetsWithDup(nums, counter, index+1, subset, res)
}
