package LeetCode

import "sort"

// 128. Longest Consecutive Sequence
// 数组无序，要求时间复杂度O(n)
func longestConsecutive(nums []int) int {
	longest := 0
	counter := make(map[int]bool)
	for _, num := range nums {
		counter[num] = false
	}
	for _, num := range nums {
		if counter[num] {
			// 已经处理过，跳过
			continue
		}
		length := 1
		counter[num] = true
		for i := num + 1; ; i++ {
			_, ok := counter[i]
			if !ok {
				break
			}
			counter[i] = true
			length++
		}
		for i := num - 1; ; i-- {
			_, ok := counter[i]
			if !ok {
				break
			}
			counter[i] = true
			length++
		}
		longest = max(length, longest)
	}
	return longest
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	// nums = [100,4,200,1,3,2]
	res, numMap := 0, map[int]int{}
	for _, num := range nums {
		if numMap[num] == 0 {
			left, right, sum := 0, 0, 0
			if numMap[num-1] > 0 {
				left = numMap[num-1]
			} else {
				left = 0
			}
			if numMap[num+1] > 0 {
				right = numMap[num+1]
			} else {
				right = 0
			}
			// sum: length of the sequence n is in
			sum = left + right + 1
			numMap[num] = sum
			// keep track of the max length
			res = max(res, sum)
			// extend the length to the boundary(s) of the sequence
			// will do nothing if n has no neighbors
			numMap[num-left] = sum
			numMap[num+right] = sum
		} else {
			continue
		}
	}
	return res
}

// 时间复杂度因为涉及了排序，至少是O(logn)
// 但实际跑起来却比上面方法快很多
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	seq, maxSeq := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			// 有重复数字，跳过
			continue
		}
		if nums[i+1]-nums[i] == 1 {
			// 有连续数字
			seq++
			if maxSeq < seq {
				maxSeq = seq
			}
			continue
		}
		seq = 1
	}
	return maxSeq
}
