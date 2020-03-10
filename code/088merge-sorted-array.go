/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (46.62%)
 * Likes:    436
 * Dislikes: 0
 * Total Accepted:    114.7K
 * Total Submissions: 245K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 * 
 * 说明:
 * 
 * 
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 * 
 * 
 * 示例:
 * 
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * 
 * 输出: [1,2,2,3,5,6]
 * 
 */

// @lc code=start
/*双指针法*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	idx1 := m - 1
	idx2 := n - 1

	idx := m + n - 1
	for idx1 >= 0 && idx2 >= 0 {
		if nums1[idx1] < nums2[idx2] {
			nums1[idx] = nums2[idx2]
			idx2--
		} else {
			nums1[idx] = nums1[idx1]
			idx1--
		}
		idx--
	}

	if idx1 < 0 {
		for i := idx2; i >= 0; i-- {
			nums1[i] = nums2[i]
		}
	}
}
// @lc code=end

