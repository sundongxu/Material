package LeetCode

/**
 * 137. Single Number II
 * 描述：
 * 难度：Medium
 * 类型：Array
 */
// 一个数出现3次，那么它各个二进制位上的1的个数也会是3的倍数
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		// i代表第i位(bit)
		total := int32(0)
		for _, num := range nums {
			// 遍历数组全体元素
			// 对数据所有元素的第i位上的1的个数进行累加，理论上应该是3的倍数+单独元素在该位的1的个数
			// 那么第i位上的1的总数total对3取模的结果，应该就是0或1，其中0说明单独元素在第i位上也是0，1则说明单独元素在第i位上是1
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			// 说明单独元素在第i位是1，那么算进结果ans中
			// ans |= 1 << i，也可以
			ans += 1 << i
		}
	}
	return int(ans)
}
