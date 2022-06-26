package LeetCode

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	// bb
	for i, _ := range s {
		if i+1 >= len(s) {
			if len(result) == 0 {
				result = string(s[i])
			}
			continue
		}
		tmp := ""
		left := i - 1
		right := i + 1
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				if len(tmp) > len(result) {
					result = tmp
				}
				break
			}
			tmp = s[left : right+1]
			left--
			right++
		}
		left = i
		right = i + 1
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				if len(tmp) > len(result) {
					result = tmp
				}
				break
			}
			tmp = s[left : right+1]
			left--
			right++
		}
	}
	return result
}
