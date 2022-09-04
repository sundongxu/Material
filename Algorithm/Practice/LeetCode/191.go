package LeetCode

/**
 * 191. Number of 1 Bits
 * 描述：
 * 难度：Easy
 * 类型：Math
 */
func hammingWeight(num uint32) int {
	cnt, bitCnt := uint32(0), 0
	for bitCnt < 32 {
		cnt += num & 1
		num >>= 1
		bitCnt++
	}
	return int(cnt)
}

// n & (n-1) => 清除最右位的二进制位1
func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
