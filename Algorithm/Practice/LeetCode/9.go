package LeetCode

import "strconv"

/**
 * 9. Palindrome Number
 * 描述：
 * 难度：Easy
 * 类型：Math
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	left := 0
	right := len(s)
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
