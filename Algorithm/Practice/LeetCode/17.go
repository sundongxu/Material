package LeetCode

/**
 * 17. Letter Combinations of a Phone Number
 * 描述：
 * 难度：Medium
 * 类型：DFS
 */
var digitLetterMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	findLetterCombinations(digits, 0, "", &res)
	return res
}

func findLetterCombinations(digits string, index int, str string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, str)
		return
	}
	options := digitLetterMap[string(digits[index])]
	for _, option := range options {
		findLetterCombinations(digits, index+1, str+string(option), res)
		// below also works, and more like backtracing
		//str := str + string(option)
		//findLetterCombinations(digits, index+1, str, res)
		//str = str[:len(str) - 1]
	}
}
