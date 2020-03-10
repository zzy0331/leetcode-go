/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (48.57%)
 * Likes:    901
 * Dislikes: 0
 * Total Accepted:    59.2K
 * Total Submissions: 120.8K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */

// @lc code=start
/*暴力解法,O(n2)*/
func trap(height []int) int {
	var maxwater int
	for i := 1; i < len(height) - 1; i++ {
		left := i
		right := i
		var leftmax, rightmax int
		for left >= 0 {
			if height[left] > height[i] && height[left] > leftmax {
				leftmax = height[left]
			}
			left--
		}

		for right <= len(height) - 1 {
			if height[right] > height[i] && height[right] > rightmax {
				rightmax = height[right]
			}
			right++
		}

		if leftmax > height[i] && rightmax > height[i] {
			if leftmax > rightmax {
				maxwater = maxwater + rightmax - height[i]
			} else {
				maxwater = maxwater + leftmax - height[i]
			}
		}
	}

	return maxwater
}

/*双指针法,O(n)*/
func trap(height []int) int {
	var maxwater int
	var leftmax, rightmax int
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftmax {
				leftmax = height[left]
			} else {
				maxwater = maxwater + leftmax - height[left]
			}
			left++
		} else {
			if height[right] > rightmax {
				rightmax = height[right]
			} else {
				maxwater = maxwater + rightmax - height[right]
			}
			right--
		}
	}

	return maxwater
}

/*用栈,看不太懂*/
func trap(height []int) int {
	var stack []int
	var maxwater int

	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			distance := i - stack[len(stack)-1] - 1
			if height[i] < height[stack[len(stack)-1]] {
				maxwater = maxwater + distance * (height[i] - height[top])
			} else {
				maxwater = maxwater + distance * (height[stack[len(stack)-1]] - height[top])
			}
		}
		stack = append(stack, i)
	}
	return maxwater
}
// @lc code=end

