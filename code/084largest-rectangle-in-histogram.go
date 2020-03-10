/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (38.70%)
 * Likes:    461
 * Dislikes: 0
 * Total Accepted:    31.1K
 * Total Submissions: 79.4K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 * 
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 * 
 * 
 * 
 * 
 * 
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 * 
 * 
 * 
 * 
 * 
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 * 
 */

// @lc code=start
/*暴力法,O(n2)*/
func largestRectangleArea(heights []int) int {
	var max int
	for i := 0; i < len(heights); i++ {
		left := i
		right := i

		for left >= 0 && heights[left] >= heights[i] {
			left--
		}

		for right <= len(heights) - 1 && heights[right] >= heights[i] {
			right++
		}

		if max < heights[i] * (right - left - 1) {
			max = heights[i] * (right - left - 1)
		}
	}

	return max
}


/*顺序栈解,O(n)*/
func largestRectangleArea(heights []int) int {
	
	stack := make([]int, 1, len(heights)+1)
	stack[0] = -1

	var max int
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack) - 1]] >= heights[i] {
			if max < heights[stack[len(stack) - 1]] * (i - stack[len(stack) - 2] - 1) {
				max = heights[stack[len(stack) - 1]] * (i - stack[len(stack) - 2] - 1)
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for len(stack) > 1 {
		if max < heights[stack[len(stack) - 1]] * (len(heights) - stack[len(stack) - 2] - 1) {
			max = heights[stack[len(stack) - 1]] * (len(heights) - stack[len(stack) - 2] - 1)
		}
		stack = stack[:len(stack)-1]
	}

	return max
}
// @lc code=end

