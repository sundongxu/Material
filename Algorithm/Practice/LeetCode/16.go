package LeetCode

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func ThreeSumClosest(nums []int, target int) int {
	result := 0
	if len(nums) <= 2 {
		return result
	}
	sort.Ints(nums)
	minGap := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			gap := abs(sum - target)
			if gap < minGap {
				result = sum
				minGap = gap
			}

			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return result
			}
		}
	}
	return result
}
