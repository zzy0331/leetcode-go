/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (59.70%)
 * Likes:    869
 * Dislikes: 0
 * Total Accepted:    192K
 * Total Submissions: 320K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 
 * 示例：
 * 
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*两个指针分别移动*/
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := &ListNode{}
	prev := dump

	cur1 := l1
	cur2 := l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			prev.Next = cur1
			cur1 = cur1.Next
		} else {
			prev.Next = cur2
			cur2 = cur2.Next
		}
		prev = prev.Next
	}

	if cur1 != nil {
		prev.Next = cur1
	} 

	if cur2 != nil {
		prev.Next = cur2
	} 

	return dump.Next
}
// @lc code=end

