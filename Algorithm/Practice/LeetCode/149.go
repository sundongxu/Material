package LeetCode

// 149. Max Points on a Line
// 时间复杂度：O(n^2 * logm)其中n为点的数量，m为横纵坐标差的最大值
// 最坏情况下我们需要枚举所有n个点
// 枚举单个点过程中需要进行O(n)次最大公约数计算
// 单次最大公约数计算的时间复杂度是O(logm)
// 因此总时间复杂度为 O(n^2 * logm)
// 空间复杂度：O(n)，其中 nn 为点的数量。主要为哈希表的开销。
func maxPoints(points [][]int) (ans int) {
	n := len(points)
	if n <= 2 {
		return n
	}
	for i, p := range points {
		if ans >= n-i || ans > n/2 {
			break
		}
		cnt := map[int]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*20001]++
		}
		for _, c := range cnt {
			ans = max(ans, c+1)
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
