package LeetCode

/**
 * 57. Insert Interval
 * 描述：
 * 难度：Medium
 * 类型：Array Interval
 */
// 区间数组已经按区间起点大小升序排列
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex][1] < newInterval[0] {
		res = append(res, intervals[curIndex]) // 持续将原有的位于新区间的左边的区间给加入结果集
		curIndex++
	}
	// curIndex此时应该指向第一个终点小于新区间起点的区间，即第一个与新区间重合的区间
	for curIndex < len(intervals) && intervals[curIndex][0] <= newInterval[1] {
		// 本循环可能执行多次，因为用当前区间的信息修改了新区间起点和终点，可能又和下一个区间重合了，所以要循环处理
		newInterval[0] = min(intervals[curIndex][0], newInterval[0])
		newInterval[1] = max(intervals[curIndex][1], newInterval[1])
		curIndex++
	}
	res = append(res, newInterval)
	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
