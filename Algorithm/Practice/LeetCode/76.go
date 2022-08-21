package LeetCode

// 76. Minimum Window Substring

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	// 分别记录t和s中某个字符出现的次数
	var tFreq, sFreq [58]int // 'A'=65，'z'=122，'z'-'A'=57，即可能访问到的下标范围为[0,57]，所以数组大小为58，其实用map更好
	// count表示窗口中的有效字符数，所谓有效字符即s当前窗口中，包含在t中，且在s当前窗口中出现次数并未超过其在t中出现次数的字符
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0

	// 统计t中各个字符出现的字符
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'A']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			// s右指针右移一位后仍未越界(当前才可继续右移)，且t中字符也未全部包含在窗口中
			sFreq[s[right+1]-'A']++ // 累加s右指针右移一位后的字符出现次数
			if sFreq[s[right+1]-'A'] <= tFreq[s[right+1]-'A'] {
				// 小于的时候累加count能理解，但为啥等于的时候还要累加呢，那s窗口中该字符的出现次数不就超过了t了么？
				// 因为上一步先累加了右移一位后对应字符的计数，所以这里小于等于的时候是应该计数的
				// 大于的时候说明某个字符重复计数了
				count++
			}
			right++
		} else {
			// s右指针已经指向s最后一个字符，即无法再右移了(否则越界) 或 有效字符计数刚好满了
			if count == len(t) && right-left+1 < minW {
				// 若此时计数值刚好与t长度相等，即意味着t中全部字符被全部包含在s当前窗口中，满足题意
				// 若当前窗口小于之前记录的最小窗口，则重新给最小窗口及其左右端点下标赋值
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if right+1 == len(s) && count < len(t) {
				// 右指针已经达到边界，并且此时有效字符计数值仍不满足要求
				// 而因为此后只能通过右移左指针缩小窗口，字符只会越来越少，不可能再满足题目要求的解了，此时可提前结束循环
				break
			}
			// 下一步准备移动s左指针，所以要把s左指针对应字符计数给相应减去
			if sFreq[s[left]-'A'] == tFreq[s[left]-'A'] {
				// s和t对于同一个字符计数相等时，若右移s左指针，则会导致窗口中包含的有效字符少了一个，所以需要将计数器减1
				// 如果sFreq值>tFreq值，则说明sFreq对当前窗口左端点字符多计数了，所以左指针右移，窗口中仍然有足够的该字符，无需将计数器递减
				// 如果sFreq值<tFreq值，则说明count < len(t)，否则至少是大于等于，此时说明s右指针已经指向最后一个字符，其实说明此时已经不可能有足够的字符了
				count--
			}
			sFreq[s[left]-'A']--
			left++
		}
	}
	if finalLeft != -1 {
		result = s[finalLeft : finalRight+1]
	}
	return result
}
