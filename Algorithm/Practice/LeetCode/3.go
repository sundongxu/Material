package LeetCode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := 0
	// abcabcbb
	for i, c := range s {
		m := make(map[rune]bool)
		m[c] = true
		// starts with c
		for _, cc := range s[i+1:] {
			_, ok := m[cc]
			if !ok {
				m[cc] = true
			} else {
				break
			}
		}
		if len(m) > result {
			result = len(m)
		}
	}
	return result
}
