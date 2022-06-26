package LeetCode

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			res[i+1] = digits[i] + 1
		} else {
			res[i+1] = res[i+1] + digits[i]
		}
		if res[i+1] >= 10 {
			res[i+1] = res[i+1] % 10
			res[i] = 1
		}
	}
	if res[0] == 0 {
		return res[1:]
	}
	return res
}
