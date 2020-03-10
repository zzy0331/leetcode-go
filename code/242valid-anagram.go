/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (58.43%)
 * Likes:    162
 * Dislikes: 0
 * Total Accepted:    79.3K
 * Total Submissions: 135K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 * 
 * 示例 1:
 * 
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: s = "rat", t = "car"
 * 输出: false
 * 
 * 说明:
 * 你可以假设字符串只包含小写字母。
 * 
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 * 
 */
// @lc code=start
//暴力：sort后判断是否相等,nlogn
//hash：记录字母出现的次数
func isAnagram(s string, t string) bool {
	dict := make(map[byte]int)

	for _, v := range s {
		_, ok := dict[byte(v)]
		if ok {
			dict[byte(v)]++
		} else {
			dict[byte(v)] = 1
		}
	}

	for _, v := range t {
		value, ok := dict[byte(v)]
		if ok && value >= 1 {
			dict[byte(v)]--
		} else {
			return false
		}
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}
// @lc code=end

