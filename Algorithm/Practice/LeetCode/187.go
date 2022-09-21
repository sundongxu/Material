package LeetCode

/**
 * 187. Repeated DNA Sequences
 * 描述：
 * 难度：Medium
 * 类型：String & Hash Map
 */
// 用一个哈希表统计s所有长度为10的子串的出现次数，返回所有出现次数超过1次的子串
// 代码实现时，可以一边遍历子串一边记录答案，为了不重复记录答案，我们只在子串出现次数到达2时记录它

const L = 10

func findRepeatedDnaSequences(s string) (ans []string) {
	cnt := map[string]int{}
	for i := 0; i <= len(s)-L; i++ {
		sub := s[i : i+L]
		cnt[sub]++
		if cnt[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return
}

// 滑动窗口
const L = 10

var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences(s string) (ans []string) {
	n := len(s)
	if n <= L {
		return
	}
	x := 0
	for _, ch := range s[:L-1] {
		x = x<<2 | bin[byte(ch)]
	}
	cnt := map[int]int{}
	for i := 0; i <= n-L; i++ {
		x = (x<<2 | bin[s[i+L-1]]) & (1<<(L*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+L])
		}
	}
	return ans
}
