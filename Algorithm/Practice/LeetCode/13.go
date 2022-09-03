package LeetCode

/**
 * 13. Roman to Integer
 * 描述：
 * 难度：Easy
 * 类型：String & Math
 */
var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// 思路：从前往后扫描，如果发现当前元素比前一个元素大，说明需要用当前元素减去前一个元素
func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastNum, res := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		num = roman[string(ch)]
		if num < lastNum {
			res -= num
		} else {
			res += num
		}
		lastNum = num
	}
	return res
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && roman[string(s[i])] > roman[string(s[i-1])] {
			res += roman[string(s[i])] - 2*roman[string(s[i-1])]
		} else {
			res += roman[string(s[i])]
		}
	}
	return res
}
