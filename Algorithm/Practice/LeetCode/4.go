package LeetCode

// 4. Median of Two Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n
	if total%2 == 0 {
		// 偶数个
		return float64(findKth(nums1, nums2, m, n, total/2)+findKth(nums1, nums2, m, n, total/2+1)) / 2.0
	}
	// 奇数个
	return float64(findKth(nums1, nums2, m, n, total/2+1))
}

func findKth(nums1, nums2 []int, m, n, k int) int {
	if m > n {
		// 总是假设 m <= n
		return findKth(nums2, nums1, n, m, k)
	}
	if m == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	ai := min(k/2, m) // 假设ai=k/2，那么bi=k/2，且k为偶数，则比较nums1[k/2 - 1]和nums2[k/2 - 1]，每次能淘汰至少k/2个数
	bi := k - ai
	if nums1[ai-1] < nums2[bi-1] {
		return findKth(nums1[ai:], nums2, m-ai, n, k-ai)
	} else if nums1[ai-1] > nums2[bi-1] {
		return findKth(nums1, nums2[bi:], m, n-bi, k-bi)
	} else {
		return nums1[ai-1]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
