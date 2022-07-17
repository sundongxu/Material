package LeetCode

import "sort"

// 49. Group Anagrams
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// 思路：将每个字符串都排序，这样同构异形词(anagram)的排序后结果都一样，用map记录，key为排序后结果，value为同构异形词数组
func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
