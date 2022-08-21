package LeetCode

// 127. Word Ladder
// 广度优先搜索BFS，按层次遍历所有可能，每层多个元素，需要把当前层所有元素遍历完毕后才能推进到下一层，类似树层次遍历
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, queue, depth := getWordMap(wordList, beginWord), []string{beginWord}, 0
	for len(queue) > 0 {
		depth++ // 本质是计数，即多往下一层，深度加一
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			word := queue[0]
			queue = queue[1:]
			candidates := getCandidates(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						// 当前候选者列表中已经包含结束词，计数加一（加上最后一个词），直接返回
						return depth + 1
					}
					delete(wordMap, candidate)
					queue = append(queue, candidate)
				}
			}
		}
	}
	return 0
}

// wordMap记录了wordList中各元素的下标
// wordMap应该排除beginWord
func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}
	return wordMap
}

// 生成并获取和word只有一字母之差的候选单词列表
func getCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(int('a')+i)+word[j+1:])
			}
		}
	}
	return res
}
