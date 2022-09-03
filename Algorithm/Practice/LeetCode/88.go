package LeetCode

/**
 * 88. Merge Sorted Array
 * 描述：
 * 难度：Easy
 * 类型：Array
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, icur := m-1, n-1, m+n-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[icur] = nums1[i1]
			i1--
		} else {
			nums1[icur] = nums2[i2]
			i2--
		}
		icur--
	}
	// 不需要，只剩nums1还有元素每用完了，但是nums1本身是有序的，说明nums1此时已经全局有序了
	//for i1 >= 0 {
	//	nums1[icur] = nums1[i1]
	//	i1--
	//	icur--
	//}
	for i2 >= 0 {
		nums1[icur] = nums2[i2]
		i2--
		icur--
	}
}
