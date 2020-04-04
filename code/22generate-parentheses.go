/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (73.77%)
 * Likes:    833
 * Dislikes: 0
 * Total Accepted:    93K
 * Total Submissions: 126K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 * 
 * 例如，给出 n = 3，生成结果为：
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 * 
 * 
 */

// @lc code=start

//left随时加，只要不超标
//right左个数>右个数
func generate(left int, right int, n int, s string, res *[]string) {
	//terminator
	if left == n && right == n {
		*res = append(*res, s)
		return
	}

	//process current logic: left, right

	//drill down
	if left < n {
		generate(left + 1, right, n, s + "(", res)
	}
	
	if left > right {
		generate(left, right + 1, n, s + ")", res)
	}
	
	//reverse states

}

func generateParenthesis(n int) []string {
	var res []string
	generate(0, 0, n, "", &res)
	return res
}
// @lc code=end

