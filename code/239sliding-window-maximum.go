/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (43.80%)
 * Likes:    252
 * Dislikes: 0
 * Total Accepted:    32K
 * Total Submissions: 72K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 * 
 * 返回滑动窗口中的最大值。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7] 
 * 解释: 
 * 
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 * 
 * 
 * 
 * 提示：
 * 
 * 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
 * 
 * 
 * 
 * 进阶：
 * 
 * 你能在线性时间复杂度内解决此题吗？
 * 
 */

// @lc code=start
/*暴力,O(n*k)*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}
	res := make([]int, len(nums) - k + 1)

	for i := 0; i < len(nums) - k + 1; i++ {
		res[i] = nums[i]
		for j :=  i + 1; j < i + k; j++ {
			if res[i] < nums[j] {
				res[i] = nums[j]
			}
		}
	}

	return res
}

	/*双端队列法*/
func maxSlidingWindow(nums []int, k int) []int {
	var res, queue []int

	for i, v := range nums {

		//向右移动窗口
		if i >= k && queue[0] <= i - k {
			queue = queue[1:]
		}

		//过滤窗口内的最大值
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= v {
			queue = queue[:len(queue)-1]
		}

		//把当前i加入到队列
		queue = append(queue, i)

		//队列最左侧即为最大值
		if i >= k - 1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

// @lc code=end

