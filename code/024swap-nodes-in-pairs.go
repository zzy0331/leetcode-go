/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (64.30%)
 * Likes:    423
 * Dislikes: 0
 * Total Accepted:    74K
 * Total Submissions: 114.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 * 
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 
 * 
 * 
 * 示例:
 * 
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
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
 /*多练习,熟能生巧,指针节点的问题都是考察指针替代关系*/
 func swapPairs(head *ListNode) *ListNode {
	prev := &ListNode{Next:head}
	pHead := prev
	
	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := prev.Next.Next

		prev.Next = b
		a.Next = b.Next
		b.Next = a
		prev = a
	}

	return pHead.Next
}
// @lc code=end

