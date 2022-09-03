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
