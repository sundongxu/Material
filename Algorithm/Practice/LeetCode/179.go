package LeetCode

import (
	"sort"
	"strconv"
)

/**
 * 179. Largest Number
 * 描述：
 * 难度：Medium
 * 类型：Array & Math
 */
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j] // 10, 100
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10 // sx = 100
		}
		for sy <= y {
			sy *= 10 // sy = 1000
		}
		return sy*x+y > sx*y+x // 1000 * 10 + 100 = 10100 100 * 100 + 10 = 10010
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := make([]byte, 0)
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
