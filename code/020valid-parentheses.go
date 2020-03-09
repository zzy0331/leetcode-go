/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.86%)
 * Likes:    1412
 * Dislikes: 0
 * Total Accepted:    213.5K
 * Total Submissions: 521.1K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 注意空字符串可被认为是有效字符串。
 * 
 * 示例 1:
 * 
 * 输入: "()"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "()[]{}"
 * 输出: true
 * 
 * 
 * 示例 3:
 * 
 * 输入: "(]"
 * 输出: false
 * 
 * 
 * 示例 4:
 * 
 * 输入: "([)]"
 * 输出: false
 * 
 * 
 * 示例 5:
 * 
 * 输入: "{[]}"
 * 输出: true
 * 
 */

// @lc code=start
/*用栈来计算匹配的左右括号*/
func isValid(s string) bool {

	/*用hash表来快速定位左右括号,否则就要写多个if else*/
	hashTbl := map[byte]byte{
		']':'[',
		')':'(',
		'}':'{',
	}

	var stack []byte
	for _, v := range s {
		value, ok := hashTbl[byte(v)]
		if ok == false {
			stack = append(stack, byte(v))
		} else {
			stackLen := len(stack)
			if stackLen == 0 || stack[stackLen-1] != value {
				return false
			}
			stack = stack[:stackLen-1]
		}
	}

	return len(stack) == 0
}
// @lc code=end

