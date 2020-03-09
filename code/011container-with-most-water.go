/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (61.33%)
 * Likes:    1128
 * Dislikes: 0
 * Total Accepted:    140.2K
 * Total Submissions: 228.5K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 * 
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 * 
 * 
 * 
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 * 
 */

// @lc code=start
// 1.枚举：left bar, right bar, (x-y)*height_diff ->O(n2)
func maxArea(height []int) int {
	/*暴力求解,O(n2)*/
	var maxAr int
	//下面两个循环是重点，可以避免循环到重复的数字
	for i := 0; i < len(height) - 1; i++ {
		for j := i + 1; j < len(height); j++ {
			var area int
			{
				if height[j] <= height[i] {
					area = (j - i) * height[j]
				} else {
					area = (j - i) * height[i]
				}
			}

			if area > maxAr {
				maxAr = area
			}
		}
	}

	return maxAr
}

/*双指针法,O(n)*/
func maxArea(height []int) int {
	var maxAr int
	left := 0
	right := len(height) - 1
	for left < right {
		var high int
		/*哪边的高度小,就把哪边的指针往中间移动*/
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		if maxAr < high * (right - left + 1) {
			maxAr = high * (right - left + 1)
		}
	}

	return maxAr
}



