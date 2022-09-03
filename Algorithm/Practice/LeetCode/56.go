package LeetCode

/**
 * 56. Merge Intervals
 * 描述：
 * 难度：Medium
 * 类型：Array Interval
 */
// 思路：先按照区间起点进行排序。然后从区间起点小的开始扫描，依次合并每个有重叠的区间。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先将全部区间按左边界大小进行排队
	quickSort(intervals, 0, len(intervals)-1)
	res := make([][]int, 0)
	res = append(res, intervals[0])
	curIndex := 0 // res下标，始终指向res的最后一个元素
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[curIndex][1] {
			// 当前区间的左边界大于已有区间的右边界，说明完全不重合，直接加入已有区间集合
			curIndex++
			res = append(res, intervals[i])
		} else {
			// 当前区间于已有区间重合，那么尝试合并，合并后的区间左边界是已有区间中末尾区间（起点最大）的左边界，右边界是前者与当前区间的右边界的最大值
			res[curIndex][1] = max(intervals[i][1], res[curIndex][1])
		}
	}
	return res
}

func partition(intervals [][]int, low, high int) int {
	pivot := intervals[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (intervals[j][0] < pivot[0]) || (intervals[j][0] == pivot[0] && intervals[j][1] < pivot[1]) {
			i++
			intervals[j], intervals[i] = intervals[i], intervals[j]
		}
	}
	intervals[i+1], intervals[high] = intervals[high], intervals[i+1]
	return i + 1
}

func quickSort(intervals [][]int, low, high int) {
	if low >= high {
		return
	}
	pivot := partition(intervals, low, high)
	quickSort(intervals, low, pivot-1)
	quickSort(intervals, pivot+1, high)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
