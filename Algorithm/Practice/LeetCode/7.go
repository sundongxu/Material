package LeetCode

import (
	"math"
	"strconv"
)

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Reverse(x int) int {
	symbol := x >= 0
	if !symbol {
		x = -x
	}
	y := strconv.Itoa(x)
	y = reverse(y)
	yy, _ := strconv.Atoi(y)
	if yy > math.MaxInt32 {
		return 0
	}
	if !symbol {
		return -yy
	}
	return yy
}
