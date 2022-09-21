package LeetCode

/**
 * 42. Trapping Rain Water
 * 描述：
 * 难度：Hard
 * 类型：
 */
func trap(height []int) int {
	// [0,1,0,2,1,0,1,3,2,1,2,1]
	volume, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				volume += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				volume += maxRight - height[right]
			}
			right--
		}
	}
	return volume
}

// 动态规划
// 创建两个长度为n的数组leftMax和rightMax
// 对于0≤i<n，leftMax[i] 表示下标i及其左边的位置中，height的最大高度，rightMax[i] 表示下标i及其右边的位置中，height的最大高度
// 显然，leftMax[0]=height[0]，rightMax[n−1]=height[n−1]
// 两个数组的其余元素的计算如下：
// 当1≤i≤n−1时，leftMax[i]=max(leftMax[i−1],height[i])
// 当0≤i≤n−2时，rightMax[i]=max(rightMax[i+1],height[i])
// 因此可以正向遍历数组height 得到数组leftMax的每个元素值，反向遍历数组height得到数组rightMax 的每个元素值
// 在得到数组leftMax和rightMax的每个元素值之后，对于0≤i<n，下标i处能接的雨水量等于min(leftMax[i],rightMax[i])−height[i]
// 遍历每个下标位置即可得到能接的雨水总量
func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
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
