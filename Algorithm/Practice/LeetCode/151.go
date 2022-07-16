package LeetCode

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	return strings.Join(reverseArray(words), " ")
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

// 思路：先反转单个单词全部字符的顺序，再反转整个句子全部字符的顺序，负负得正，最终得到反转了全部单词顺序的句子（字符串）
func reverseWords(s string) string {
	cArr := []rune(s)
	i, j, k := 0, 0, 0
	wordCnt := 0
	for {
		for i < len(cArr) && cArr[i] == ' ' {
			i++
		}
		if i == len(cArr) {
			break
		}
		if wordCnt > 0 {
			cArr[j] = ' '
			j++
		}
		k = j
		for i < len(cArr) && cArr[i] != ' ' {
			cArr[j] = cArr[i]
			i++
			j++
		}
		// 此时单词的左边界是一开始记录j值的k，右边界是j-1
		reverseArr(cArr, k, j-1)
		wordCnt++
	}
	// 单词的长度一定<=原字符串，因为反转过程中会有多余空格被去掉
	// 此时整个字符串的左边界为0，右边界是j-1
	reverseArr(cArr, 0, j-1)
	return string(cArr[:j])
}

func reverseArr(cArr []rune, start, end int) string {
	for i, j := start, end; i < j; {
		cArr[i], cArr[j] = cArr[j], cArr[i]
		i++
		j--
	}
	return string(cArr)
}
