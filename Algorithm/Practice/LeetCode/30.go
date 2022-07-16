package LeetCode

// 30. Substring with Concatenation of All Words

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	totalWordLen := wordLen * len(words)
	if len(s) < totalWordLen {
		return []int{}
	}
	res := make([]int, 0)
	wordCnt := make(map[string]int)
	for _, word := range words {
		// 可能有重复
		wordCnt[word]++
	}
	for i := 0; i <= len(s)-totalWordLen; i++ {
		// i为父串起始下标
		unusedWordCnt := deepCopy(wordCnt) // map一定要深拷贝，直接赋值是浅拷贝指向同一块内存地址
		for j := i; j != i+totalWordLen; j += wordLen {
			// j为父串中用于匹配子串的起始下标
			word := s[j : j+wordLen]     // 从s中截取长度为wordLen的子串，认为是一个单词
			_, ok := unusedWordCnt[word] // 在未使用单词字典中查找该单词
			if ok {
				// 字典中存在，说明未使用或未使用完，那么递减其计数表示使用该单词，如计数本次减为0，那么从该字典中删除该单词，表示其已全部使用
				unusedWordCnt[word]--
				if unusedWordCnt[word] == 0 {
					delete(unusedWordCnt, word)
				}
				// 不跳出循环，继续查找下一个单词
			} else {
				// 字典中不存在，说明已使用，那么跳出循环，递增父串起始下标i，重新开始查找单词
				break
			}
		}
		if len(unusedWordCnt) == 0 {
			res = append(res, i)
		}
	}
	return res
}

func deepCopy(m map[string]int) map[string]int {
	nm := make(map[string]int)
	for k, v := range m {
		nm[k] = v
	}
	return nm
}
