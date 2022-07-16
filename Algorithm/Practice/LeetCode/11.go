package LeetCode

func MaxArea(height []int) int {
	result := 0
	minHeight := 0
	for i := 0; i < len(height); i++ {
		if height[i] < minHeight {
			continue
		}
		for j := len(height) - 1; j >= 0; j-- {
			if height[j] < minHeight {
				continue
			}
			area := 0
			tmpMinHeight := 0
			if height[i] < height[j] {
				area = (j - i) * height[i]
				tmpMinHeight = height[i]
			} else {
				area = (j - i) * height[j]
				tmpMinHeight = height[j]
			}
			if area > result {
				result = area
				minHeight = tmpMinHeight
			}
		}
	}
	return result
}

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
