package LeetCode

/**
 * 190. Reverse Bits
 * 描述：
 * 难度：Easy
 * 类型：Math
 */
func reverseBits(num uint32) uint32 {
	res, cnt := uint32(0), 0
	for cnt < 32 {
		lastBit := num & 1
		res = (res << 1) + lastBit // res = (res << 1) | (num & 1)
		num = num >> 1             // num >>= 1
		cnt++
	}
	return res
}
