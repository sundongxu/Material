package LeetCode

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	numPerGroup := 2*numRows - 2
	groupNum := len(s) / (2*numRows - 2)
	result := make([]byte, 0)
	// i = 0
	tmp := 0
	for tmp < len(s) {
		result = append(result, s[tmp])
		tmp += numPerGroup
	}
	for i := 1; i < numRows-1; i++ {
		// by row
		for j := 1; j <= groupNum; j++ {
			// by group
			// row(i) group(n): i-1+(n-1)*numPerGroup, numPerGroup-(i-1)+(n-1)*numPerGroup
			result = append(result, s[i+(j-1)*numPerGroup], s[numPerGroup-i+(j-1)*numPerGroup])
		}
		if groupNum*numPerGroup < len(s) {
			if i+groupNum*numPerGroup < len(s) {
				result = append(result, s[i+groupNum*numPerGroup])
			}
			if numPerGroup-i+groupNum*numPerGroup < len(s) {
				result = append(result, s[numPerGroup-i+groupNum*numPerGroup])
			}
		}
	}
	// i = numRows-1
	tmp = numRows - 1
	for tmp < len(s) {
		result = append(result, s[tmp])
		tmp += numPerGroup
	}
	// numRows = 4, numPerGroup = 6
	// row(1): 0 numPerGroup numPerGroup+numPerGroup numPerGroup+numPerGroup+numPerGroup
	// row(2): 1 numPerGroup-1 1+numPerGroup  numPerGroup-1+numPerGroup... =>  1,5,7,11,13,17
	// row(n): n-1 numPerGroup-(n-1) numPerGroup+(n-1) numPerGroup-(n-1)+numPerGroup
	// row(numRows): numRows-1 numRow-1+numPerGroup numRow-1+numPerGroup+numPerGroup
	return string(result)
}
