/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (60.75%)
 * Likes:    282
 * Dislikes: 0
 * Total Accepted:    54.6K
 * Total Submissions: 89.4K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

// @lc code=start
/*hash解法,golang的map中key可以存数组,数组是值拷贝*/
func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	var res [][]string

	for i := 0; i < len(strs); i++ {
		var tmp [26]int
		for _, v := range strs[i] {
			tmp[v-'a']++
		}
		hash[tmp] = append(hash[tmp], strs[i])
	}

	for _, v := range hash {
		res = append(res, v)
	}

	return res
}
// @lc code=end

