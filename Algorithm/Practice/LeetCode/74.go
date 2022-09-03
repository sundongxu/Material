package LeetCode

/**
 * 74. Search a 2D Matrix
 * 描述：
 * 难度：Medium
 * 类型：Matrix
 */
// 每次都找右上角的元素，假定它不等于目标值，如果比目标值小，那么可以排除其所在行，否则，则可以排除其所在列
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := 0, len(matrix[0])-1
	for m <= len(matrix)-1 && n >= 0 {
		rightCorner := matrix[m][n]
		if rightCorner == target {
			return true
		} else if rightCorner > target {
			n--
		} else {
			m++
		}
	}
	return false
}

// 二分查找
//关于取中间数 mid = (left + right) / 2;
//在 left + right 很大的时候会发生整型溢出，一般这样改写：mid = left + (right - left) / 2
//还有一个细节，/ 2 表示的是下取整，当数组中的元素个数为偶数的时候，只能取到位于左边的那个元素
//取右边中间数的表达式是mid = left + (right - left + 1) / 2
func SearchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	// 注意这里last一定要等于元素总个数或最后一个元素的下标+1（类似C++中的vector.end()），而非最后一个元素的下标，否则无法处理当矩阵只有1行1列时的情况
	// 若last是最后一个元素的下标(rows*cols-1)，那么此时first=last=0，将无法进入二分查找循环，直接返回false，这是不对的
	// 而first表示第一个元素的下标
	first, last := 0, rows*cols
	for first < last {
		mid := first + (last-first)/2
		value := matrix[mid/cols][mid%cols]
		if value == target {
			return true
		} else if value > target {
			// mid > target，那么target元素一定在mid的左边，last = mid
			// 为什么不能last = mid - 1，因为last表示的含义就是最后一个有效元素的下一个位置，如果last=mid-1的话，就会跳过mid-1这个可能答案
			last = mid
		} else {
			// value < target
			first = mid + 1
		}
	}
	return false
}
