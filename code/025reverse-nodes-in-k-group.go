/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (56.19%)
 * Likes:    386
 * Dislikes: 0
 * Total Accepted:    39.6K
 * Total Submissions: 70.2K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * 
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * 
 * 示例 :
 * 
 * 给定这个链表：1->2->3->4->5
 * 
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * 
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 * 
 * 说明 :
 * 
 * 
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
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
 func reverseKGroup(head *ListNode, k int) *ListNode {
	dump := &ListNode{Next:head}
	prev := dump

	curr := head

	if k <= 1 {
		return head
	}

	/*先计算反转的次数*/
	var nodecnt int
	for curr != nil {
		nodecnt++
		curr = curr.Next
	}

	curr = head
	/*开始每组每组反转*/
	for i := 0; i < nodecnt/k; i++ {
		//依次把最右边节点移到最左边,该处的做法与反转链表的做法不一致，需谨记
		for j := 0; j < k - 1; j++ {
			temp := curr.Next
			curr.Next = temp.Next
			temp.Next = prev.Next
			prev.Next = temp
		}
		prev = curr
		curr = curr.Next
	}

	return dump.Next
}
// @lc code=end

